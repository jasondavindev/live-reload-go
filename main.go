package main

import (
	"log"
	"os"

	"github.com/jasondavindev/hacktoberfest-2020/domain"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	directoryWatch := os.Getenv("DIRECTORY")
	excludedDirectories := os.Getenv("EXCLUDE")

	watcher := domain.CreateWatcher()
	defer domain.CloseWatcher(watcher)

	done := make(chan bool)
	go domain.ListenEvents(watcher, excludedDirectories)

	domain.AddDirectoryToWatch(watcher, directoryWatch)
	<-done
}
