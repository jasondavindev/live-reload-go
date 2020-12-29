package command

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

type JobRunner struct {
	jobs []job
}

type job struct {
	cmdName string
	cmdArgs []string
}

func createJob(commandStr string) job {
	if commandStr == "" {
		log.Fatal("No command to execute was specified")
	}

	commands := strings.Split(commandStr, " ")
	job := job{
		cmdName: commands[0],
		cmdArgs: commands[1:],
	}

	return job
}

func (j *job) executeJob() string {
	cmd := exec.Command(j.cmdName, j.cmdArgs...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	cmd.Run()

	outStr := string(stdout.Bytes())

	return outStr
}

func CreateJobRunner(commandList []string) JobRunner {
	jobRunner := JobRunner{}
	for _, cmd := range commandList {
		jobRunner.AddJob(createJob(cmd))
	}

	return jobRunner
}

func (j *JobRunner) AddJob(job job) {
	j.jobs = append(j.jobs, job)
}

func (j *JobRunner) RunJobs() []string {
	res := []string{}

	for _, job := range j.jobs {
		res = append(res, job.executeJob())
	}

	return res
}
