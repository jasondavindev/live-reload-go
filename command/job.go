package command

import (
	"log"
	"os/exec"
	"strings"
)

type Job struct {
	cmdName string
	cmdArgs []string
}

func CreateJob(commandStr string) Job {
	if commandStr == "" {
		log.Fatal("No command to execute was specified")
	}

	commands := strings.Split(commandStr, " ")
	job := Job{
		cmdName: commands[0],
		cmdArgs: commands[1:],
	}

	return job
}

func (j *Job) ExecuteJob() string {
	cmd := exec.Command(j.cmdName, j.cmdArgs...)

	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err.Error())
	}

	return string(stdout)
}
