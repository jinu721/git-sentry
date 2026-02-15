package state

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDefaultState(t *testing.T) {
	state := DefaultState()
	
	if state.FilesChanged != 0 {
		t.Error("FilesChanged should be 0 initially")
	}
	
	if state.LinesAdded != 0 {
		t.Error("LinesAdded should be 0 initially")
	}
	
	if state.LinesRemoved != 0 {
		t.Error("LinesRemoved should be 0 initially")
	}
	
	if state.LastCommit.IsZero() == false {
		t.Error("LastCommit should be zero time initially")
	}
}

func TestStateSaveLoad(t *testing.T) {
	tempDir := "test_state"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	originalState := DefaultState()
	originalState.FilesChanged = 5
	originalState.LinesAdded = 100
	
	err := originalState.Save(tempDir)
	if err != nil {
		t.Fatalf("Failed to save state: %v", err)
	}
	
	loadedState, err := Load(tempDir)
	if err != nil {
		t.Fatalf("Failed to load state: %v", err)
	}
	
	if loadedState.FilesChanged != originalState.FilesChanged {
		t.Error("FilesChanged mismatch after save/load")
	}
	
	if loadedState.LinesAdded != originalState.LinesAdded {
		t.Error("LinesAdded mismatch after save/load")
	}
}

func TestStateThreadSafety(t *testing.T) {
	state := DefaultState()
	
	done := make(chan bool)
	
	go func() {
		for i := 0; i < 100; i++ {
			state.IncrementFilesChanged()
		}
		done <- true
	}()
	
	go func() {
		for i := 0; i < 100; i++ {
			state.GetStats()
		}
		done <- true
	}()
	
	<-done
	<-done
	
	files, _, _, _, _ := state.GetStats()
	if files != 100 {
		t.Errorf("Expected 100 files changed, got %d", files)
	}
}

func TestStateRecordCommit(t *testing.T) {
	state := DefaultState()
	state.FilesChanged = 5
	state.LinesAdded = 50
	
	state.RecordCommit()
	
	files, lines, _, lastCommit, _ := state.GetStats()
	
	if files != 0 {
		t.Error("FilesChanged should be reset after commit")
	}
	
	if lines != 0 {
		t.Error("LinesAdded should be reset after commit")
	}
	
	if lastCommit.IsZero() {
		t.Error("LastCommit should be set after commit")
	}
}

func TestStateRecordPush(t *testing.T) {
	state := DefaultState()
	
	state.RecordPush()
	
	_, _, _, _, lastPush := state.GetStats()
	
	if lastPush.IsZero() {
		t.Error("LastPush should be set after push")
	}
}