package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "no command specified: valid commands: [run]")
		os.Exit(1)
	}

	args := os.Args[1:]
	if len(args) <= 1 {
		fmt.Fprintln(os.Stderr, "no command to run")
		os.Exit(1)
	}
	if args[0] == "run" {
		cmd := exec.Command(args[1], args[2:]...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		err := cmd.Run()
		var exitError *exec.ExitError
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			exitCode := 1
			if errors.As(err, &exitError) {
				exitCode = exitError.ExitCode()
			}
			os.Exit(exitCode)
		}
	} else {
		fmt.Fprintf(os.Stderr, "invalid command: %s\n", args[0])
		os.Exit(1)
	}
}
