package state

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
	
	"gitsentry/internal/security"
)

type State struct {
	mu           sync.RWMutex
	FilesChanged   int       `json:"files_changed"`
	LinesAdded     int       `json:"lines_added"`
	LinesRemoved   int       `json:"lines_removed"`
	LastCommit     time.Time `json:"last_commit"`
	LastPush       time.Time `json:"last_push"`
	LastActivity   time.Time `json:"last_activity"`
}

func DefaultState() *State {
	return &State{
		FilesChanged:   0,
		LinesAdded:     0,
		LinesRemoved:   0,
		LastCommit:     time.Time{},
		LastPush:       time.Time{},
		LastActivity:   time.Now(),
	}
}

func Load(gitsentryDir string) (*State, error) {
	statePath := filepath.Join(gitsentryDir, "state.json")
	
	// If state doesn't exist, create default
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		state := DefaultState()
		if err := state.Save(gitsentryDir); err != nil {
			return nil, err
		}
		return state, nil
	}
	
	// Load existing state
	data, err := security.SecureReadFile(statePath)
	if err != nil {
		return nil, err
	}
	
	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	
	return &state, nil
}

func (s *State) Save(gitsentryDir string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	statePath := filepath.Join(gitsentryDir, "state.json")
	
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	
	return security.SecureWriteFile(statePath, data)
}

func (s *State) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.FilesChanged = 0
	s.LinesAdded = 0
	s.LinesRemoved = 0
}

func (s *State) RecordCommit() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.LastCommit = time.Now()
	s.FilesChanged = 0
	s.LinesAdded = 0
	s.LinesRemoved = 0
}

func (s *State) RecordPush() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.LastPush = time.Now()
}

func (s *State) IncrementFilesChanged() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.FilesChanged++
	s.LastActivity = time.Now()
}

func (s *State) GetStats() (int, int, int, time.Time, time.Time) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	return s.FilesChanged, s.LinesAdded, s.LinesRemoved, s.LastCommit, s.LastPush
}