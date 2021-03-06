package listener

import (
	"testing"

	"github.com/jasondavindev/hacktoberfest-2020/command"

	"github.com/stretchr/testify/assert"
	"gopkg.in/fsnotify.v1"
)

func TestIsModifiedFile(t *testing.T) {
	assert.Equal(t, true, isModifiedFile(fsnotify.Event{Op: fsnotify.Write}))
	assert.Equal(t, true, isModifiedFile(fsnotify.Event{Op: fsnotify.Remove}))
	assert.Equal(t, true, isModifiedFile(fsnotify.Event{Op: fsnotify.Create}))
	assert.Equal(t, false, isModifiedFile(fsnotify.Event{Op: fsnotify.Chmod}))
}

func TestIsExcludedFile(t *testing.T) {
	jr := command.CreateJobRunner([]string{"echo", "ls"})
	filepath := "/etc/file.go"
	excludedFiles := "file.go"
	listener := CreateChangesListener(excludedFiles, jr)

	assert.Equal(t, true, listener.isExcludedFile(filepath))

	listener.excludedDirectories = []string{}

	assert.Equal(t, false, listener.isExcludedFile(filepath))
}
