package domain

import (
	s "strings"
)

func commandExecutor(command string) {
	var app string
	if command != "" {
		commands := s.Split(command, " ")
		app = commands[0]
		executeCommand(app, commands[1:])
	}
}
