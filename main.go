package main

import (
	"strings"

	"github.com/jasondavindev/hacktoberfest-2020/config"
	"github.com/jasondavindev/hacktoberfest-2020/domain"
)

func main() {
	cfg := config.CfgFactory()

	directoryWatch := cfg.Directory
	excludedDirectories := strings.Split(cfg.Exclude, ",")
	command := cfg.Command

	watcher := domain.CreateWatcher()
	defer domain.CloseWatcher(watcher)

	done := make(chan bool)
	go domain.ListenEvents(watcher, excludedDirectories, command)

	domain.SetupDirectoriesToWatch(watcher, directoryWatch)
	<-done
}
