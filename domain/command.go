package domain

import (
	"log"
	"os/exec"
	s "strings"
)

// ExecuteCommand executa um comando no shell da maquina
func executeCommand(str string) {

	commands := s.Split(str, " ")

	cmd := exec.Command(commands[0], commands[1:]...)

	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Print(string(stdout))
}
