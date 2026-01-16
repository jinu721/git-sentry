package monitor

import (
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type FileMonitor struct {
	watcher  *fsnotify.Watcher
	callback func(string)
	done     chan bool
}

func NewFileMonitor(path string, callback func(string)) (*FileMonitor, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	
	monitor := &FileMonitor{
		watcher:  watcher,
		callback: callback,
		done:     make(chan bool),
	}
	
	// Add the root path
	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}
	
	// Start monitoring
	go monitor.watch()
	
	return monitor, nil
}

func (fm *FileMonitor) watch() {
	for {
		select {
		case event, ok := <-fm.watcher.Events:
			if !ok {
				return
			}
			
			// Filter out events we don't care about
			if fm.shouldIgnore(event.Name) {
				continue
			}
			
			// Only care about write and create events
			if event.Op&fsnotify.Write == fsnotify.Write || 
			   event.Op&fsnotify.Create == fsnotify.Create {
				fm.callback(event.Name)
			}
			
		case err, ok := <-fm.watcher.Errors:
			if !ok {
				return
			}
			// Log error but continue monitoring
			_ = err
			
		case <-fm.done:
			return
		}
	}
}

func (fm *FileMonitor) shouldIgnore(path string) bool {
	// Ignore hidden directories and files
	if strings.Contains(path, "/.") {
		return true
	}
	
	// Ignore common build/cache directories
	ignoreDirs := []string{
		"node_modules",
		".git",
		".gitsentry",
		"vendor",
		"target",
		"build",
		"dist",
		".cache",
		"tmp",
		"temp",
	}
	
	for _, dir := range ignoreDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}
	
	// Ignore common temporary file extensions
	ignoreExts := []string{
		".tmp",
		".temp",
		".swp",
		".swo",
		".log",
		".lock",
		"~",
	}
	
	ext := filepath.Ext(path)
	for _, ignoreExt := range ignoreExts {
		if ext == ignoreExt || strings.HasSuffix(path, ignoreExt) {
			return true
		}
	}
	
	return false
}

func (fm *FileMonitor) Stop() {
	close(fm.done)
	fm.watcher.Close()
}