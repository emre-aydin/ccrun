package main_test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os/exec"
	"testing"
)

func TestCommandLine(t *testing.T) {
	command := exec.Command("./ccrun", "run", "echo", "Hello Coding Challenges!")
	command.Dir = "/home/emre/projects/ccrun"

	stdoutPipe, err := command.StdoutPipe()
	assert.Nil(t, err)

	err = command.Start()
	assert.Nil(t, err)

	stdoutMsg, err := io.ReadAll(stdoutPipe)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Coding Challenges!\n", string(stdoutMsg))

	err = command.Wait()
	assert.Nil(t, err)
}
