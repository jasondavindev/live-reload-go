package domain

import (
	"fmt"
	"log"

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

func ListenEvents(watcher *fsnotify.Watcher, excludedDirectories string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if checkExcludedDirectories(event.Name, excludedDirectories) {
				continue
			}

			if isModifiedFile(event) {
				fmt.Println("modified file:", event.Name)
			}

			command()
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func checkExcludedDirectories(eventName string, excludedDirectories string) bool {
	return eventName == excludedDirectories
}

func isModifiedFile(e fsnotify.Event) bool {
	return e.Op == fsnotify.Create || e.Op == fsnotify.Remove
}

func AddDirectoryToWatch(w *fsnotify.Watcher, directory string) {
	err := w.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
}
