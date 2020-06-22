package main

import (
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string, argsarr []string) (err error)  {
	args := argsarr
	cmdobj := exec.Command(command)

	cmdobj.Stdout = os.Stdout
	cmdobj.Stderr = os.Stderr
	err = cmdobj.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	command :=  "date"
	executeCommand(command, argsarr []string{"-l"})
}