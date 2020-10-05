package main

import (
	"github.com/jasondavindev/hacktoberfest-2020/config"
	"github.com/jasondavindev/hacktoberfest-2020/listener"
)

func main() {

	configPath := config.CfgFilePath()

	cfg := config.CfgFactory(configPath)

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
