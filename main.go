package main

import (
	"github.com/jasondavindev/hacktoberfest-2020/domain"
)

func main() {

	cfg := CfgFactory()

	directoryWatch := cfg.Directory
	excludedDirectories := cfg.Exclude

	watcher := domain.CreateWatcher()
	defer domain.CloseWatcher(watcher)

	done := make(chan bool)
	go domain.ListenEvents(watcher, excludedDirectories)

	domain.AddDirectoryToWatch(watcher, directoryWatch)
	<-done
}
