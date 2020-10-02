package domain

import (
	"os"
)

func command() {
	env := os.Getenv("COMMAND")

	if env != "" {
		executeCommand(env)
	}
}
