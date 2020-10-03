package command

import (
	"log"
	"os/exec"
	"strings"
)

var commandName string
var commandArgs []string

func CreateCommand(commandStr string) {
	if commandStr == "" {
		log.Fatal("No command to execute was specified")
	}

	commands := strings.Split(commandStr, " ")
	commandName = commands[0]
	commandArgs = commands[1:]
}

func ExecuteJob() string {
	cmd := exec.Command(commandName, commandArgs...)

	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err.Error())
	}

	return string(stdout)
}
