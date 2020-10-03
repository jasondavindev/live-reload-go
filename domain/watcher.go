package domain

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/fsnotify.v1"
)

func CreateWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	return watcher
}

func CloseWatcher(w *fsnotify.Watcher) {
	w.Close()
}

func ListenEvents(watcher *fsnotify.Watcher, excludedDirectories []string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if isExcludedFile(event.Name, excludedDirectories) {
				continue
			}

			eventHandler(event)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			log.Println("error:", err)
		}
	}
}

func isModifiedFile(e fsnotify.Event) bool {
	return e.Op == fsnotify.Create || e.Op == fsnotify.Remove
}

func isExcludedFile(absoluteFile string, excludedFiles []string) bool {
	fileName := filepath.Base(absoluteFile)

	for _, file := range excludedFiles {
		if fileName == file {
			return true
		}
	}

	return false
}

func SetupDirectoriesToWatch(w *fsnotify.Watcher, directory string) {
	err := w.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
}

func splitExcludedFiles(excludedFiles string) []string {
	return strings.Split(excludedFiles, ",")
}

func eventHandler(event fsnotify.Event) {
	if isModifiedFile(event) {
		fmt.Println("modified file:", event.Name)
	}
}
