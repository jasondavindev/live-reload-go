package main

import (
	"github.com/jasondavindev/hacktoberfest-2020/config"
	"github.com/jasondavindev/hacktoberfest-2020/listener"
)

func main() {
	cfg := config.CfgFactory()

	directoryWatch := cfg.Directory
	excludedDirectories := cfg.Exclude
	command := cfg.Command

	cl := listener.CreateChangesListener(excludedDirectories, command)
	defer cl.CloseWatcher()

	done := make(chan bool)
	go cl.ListenEvents()

	cl.SetupDirectoriesToWatch(directoryWatch)
	<-done
}
