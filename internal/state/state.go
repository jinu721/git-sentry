package state

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type State struct {
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
	data, err := os.ReadFile(statePath)
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
	statePath := filepath.Join(gitsentryDir, "state.json")
	
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(statePath, data, 0644)
}

func (s *State) Reset() {
	s.FilesChanged = 0
	s.LinesAdded = 0
	s.LinesRemoved = 0
}

func (s *State) RecordCommit() {
	s.LastCommit = time.Now()
	s.Reset()
}

func (s *State) RecordPush() {
	s.LastPush = time.Now()
}