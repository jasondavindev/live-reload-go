package domain

import (
	"os"
	s "strings"
)

func command() {
	env := os.Getenv("COMMAND")

	var app string
	if env != "" {
		commands := s.Split(env, " ")
		app = commands[0]
		executeCommand(app, commands[1:])
	}
}
