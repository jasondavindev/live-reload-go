package listener

import (
	"testing"

	"gopkg.in/fsnotify.v1"
)

func TestIsModifiedFile(t *testing.T) {
	e := fsnotify.Event{Op: fsnotify.Write}

	if !isModifiedFile(e) {
		t.Errorf("Expected isModifiedFile %s is true", e.Op)
	}

	e.Op = fsnotify.Remove

	if !isModifiedFile(e) {
		t.Errorf("Expected isModifiedFile %s is true", e.Op)
	}

	e.Op = fsnotify.Remove

	if !isModifiedFile(e) {
		t.Errorf("Expected isModifiedFile %s is true", e.Op)
	}

	e.Op = fsnotify.Chmod

	if isModifiedFile(e) {
		t.Errorf("Expected isModifiedFile %s is false", e.Op)
	}
}

func TestIsExcludedFile(t *testing.T) {
	filepath := "/etc/file.go"
	excludedFiles := "file.go"
	listener := CreateChangeListener(excludedFiles, "echo")

	if !listener.isExcludedFile(filepath) {
		t.Errorf("Expected isExcludedFile to be true")
	}

	listener.excludedDirectories = []string{}

	if listener.isExcludedFile(filepath) {
		t.Errorf("Expected isExcludedFile to be false")
	}
}
