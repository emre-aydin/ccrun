package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const AlpineRootFs = "/home/emre/projects/ccrun/alpine"

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

	switch args[0] {
	case "run":
		// replace 'run' with 'wrap-run' to start the new process in a new UTS namespace
		args[0] = "wrap-run"
		cmd := exec.Command(os.Args[0], args...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
			UidMappings: []syscall.SysProcIDMap{
				{ContainerID: 0, HostID: 1000, Size: 1},
			},
		}
		//err := createCgroup("my-container")
		//if err != nil {
		//	exitWithError(err)
		//}
		//defer deleteCgroup("my-container")

		err := cmd.Run()
		if err != nil {
			exitWithError(err)
		}
	case "wrap-run":
		var err error

		err = syscall.Sethostname([]byte("container"))
		if err != nil {
			exitWithError(err)
		}

		err = syscall.Mount("/proc", filepath.Join(AlpineRootFs, "proc"), "proc", 0, "")
		if err != nil {
			exitWithError(err)
		}

		cmd := exec.Command(args[1], args[2:]...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Chroot: AlpineRootFs,
		}
		cmd.Dir = "/"
		err = cmd.Run()
		if err != nil {
			exitWithError(err)
		}
	default:
		fmt.Fprintf(os.Stderr, "invalid command: %s\n", args[0])
		os.Exit(1)
	}
}

func exitWithError(err error) {
	var exitError *exec.ExitError
	fmt.Fprintln(os.Stderr, err.Error())
	exitCode := 1
	if errors.As(err, &exitError) {
		exitCode = exitError.ExitCode()
	}
	os.Exit(exitCode)
}

func createCgroup(name string) error {
	err := os.Mkdir(fmt.Sprintf("/sys/fs/cgroup/%s", name), 0750)
	return err
}

func deleteCgroup(name string) error {
	return os.RemoveAll(fmt.Sprintf("/sys/fs/cgroup/%s", name))
}
