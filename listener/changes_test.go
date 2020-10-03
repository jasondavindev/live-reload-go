package listener

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jasondavindev/hacktoberfest-2020/mocks"

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
	listener := CreateChangesListener(excludedFiles, "echo")

	if !listener.isExcludedFile(filepath) {
		t.Errorf("Expected isExcludedFile to be true")
	}

	listener.excludedDirectories = []string{}

	if listener.isExcludedFile(filepath) {
		t.Errorf("Expected isExcludedFile to be false")
	}
}

func TestEventHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	// event := fsnotify.Event{Op: fsnotify.Create}

	mocks.NewMockIChangesListener(mockCtrl)
	// mockChangesListener.EXPECT().EventHandler(event)
}
