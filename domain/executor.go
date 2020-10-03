package domain

import (
	s "strings"
)

func commandExecutor(command string) {
	if command == "" {
		return
	}

	commands := s.Split(command, " ")
	app := commands[0]
	executeCommand(app, commands[1:])
}
