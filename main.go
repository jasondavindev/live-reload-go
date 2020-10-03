package main

import (
	"strings"

	cmd "github.com/jasondavindev/hacktoberfest-2020/command"
	"github.com/jasondavindev/hacktoberfest-2020/config"
	"github.com/jasondavindev/hacktoberfest-2020/listener"
)

func main() {
	cfg := config.CfgFactory()

	directoryWatch := cfg.Directory
	excludedDirectories := strings.Split(cfg.Exclude, ",")
	command := cfg.Command

	cmd.CreateCommand(command)
	watcher := listener.CreateWatcher()
	defer listener.CloseWatcher(watcher)

	done := make(chan bool)
	go listener.ListenEvents(watcher, excludedDirectories, command)

	listener.SetupDirectoriesToWatch(watcher, directoryWatch)
	<-done
}
