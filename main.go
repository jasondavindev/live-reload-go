package main

import (
	"github.com/jasondavindev/hacktoberfest-2020/command"
	"github.com/jasondavindev/hacktoberfest-2020/config"
	"github.com/jasondavindev/hacktoberfest-2020/listener"
)

func main() {

	configPath := config.CfgFilePath()

	cfg := config.CfgFactory(configPath)

	directoryWatch := cfg.Directory
	excludedDirectories := cfg.Exclude
	commands := cfg.Commands

	jr := command.CreateJobRunner(commands)
	cl := listener.CreateChangesListener(excludedDirectories, jr)
	defer cl.CloseWatcher()

	done := make(chan bool)
	go cl.ListenEvents()

	cl.SetupDirectoriesToWatch(directoryWatch)
	<-done
}
