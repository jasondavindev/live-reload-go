package domain

import (
	"log"
	"os/exec"
)

func executeCommand(app string, args []string) {

	cmd := exec.Command(app, args...)

	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Print(string(stdout))
}
