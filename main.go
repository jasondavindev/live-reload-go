package main

import (
	"github.com/jasondavindev/hacktoberfest-2020/domain"
)

func main() {

	cfg := CfgFactory()

	directoryWatch := cfg.Directory
	excludedDirectories := cfg.Exclude
	command := cfg.Command

	watcher := domain.CreateWatcher()
	defer domain.CloseWatcher(watcher)

	done := make(chan bool)
	go domain.ListenEvents(watcher, excludedDirectories, command)

	domain.AddDirectoryToWatch(watcher, directoryWatch)
	<-done
}
