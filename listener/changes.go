package listener

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/jasondavindev/hacktoberfest-2020/command"
	"gopkg.in/fsnotify.v1"
)

type ChangeListener struct {
	watcher             *fsnotify.Watcher
	excludedDirectories []string
	job                 command.Job
}

func CreateChangeListener(excludedDirectories string, cmd string) ChangeListener {
	listener := ChangeListener{}
	listener.watcher = CreateWatcher()
	listener.excludedDirectories = splitExcludedFiles(excludedDirectories)
	listener.job = command.CreateJob(cmd)

	return listener
}

func CreateWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	return watcher
}

func isModifiedFile(e fsnotify.Event) bool {
	return e.Op == fsnotify.Create || e.Op == fsnotify.Remove || e.Op == fsnotify.Write
}

func splitExcludedFiles(excludedFiles string) []string {
	return strings.Split(excludedFiles, ",")
}

func (cl *ChangeListener) CloseWatcher() {
	cl.watcher.Close()
}

func (cl *ChangeListener) ListenEvents() {
	for {
		select {
		case event, ok := <-cl.watcher.Events:
			if !ok {
				return
			}

			if cl.isExcludedFile(event.Name) {
				continue
			}

			cl.eventHandler(event)
		case err, ok := <-cl.watcher.Errors:
			if !ok {
				return
			}

			log.Println("error:", err)
		}
	}
}

func (cl *ChangeListener) isExcludedFile(absoluteFile string) bool {
	fileName := filepath.Base(absoluteFile)

	for _, file := range cl.excludedDirectories {
		if fileName == file {
			return true
		}
	}

	return false
}

func (cl *ChangeListener) SetupDirectoriesToWatch(directory string) {
	err := cl.watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
}

func (cl *ChangeListener) eventHandler(event fsnotify.Event) {
	if isModifiedFile(event) {
		fmt.Println(cl.job.ExecuteJob())
	}
}
