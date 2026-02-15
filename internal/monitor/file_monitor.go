package monitor

import (
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"gitsentry/internal/security"
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
	
	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}
	
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
			
			if fm.shouldIgnore(event.Name) {
				continue
			}
			
			if err := security.ValidateFilePath(event.Name); err != nil {
				continue
			}
			
			if event.Op&fsnotify.Write == fsnotify.Write || 
			   event.Op&fsnotify.Create == fsnotify.Create {
				fm.callback(event.Name)
			}
			
		case _, ok := <-fm.watcher.Errors:
			if !ok {
				return
			}
			
		case <-fm.done:
			return
		}
	}
}

func (fm *FileMonitor) shouldIgnore(path string) bool {
	if strings.Contains(path, "/.") {
		return true
	}
	
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