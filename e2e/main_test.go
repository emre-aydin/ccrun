package e2e_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCommandLine(t *testing.T) {
	binaryName := buildBinary(t)
	assert.NotNil(t, binaryName)
	defer func(name string) {
		err := os.Remove(fmt.Sprintf("../%s", name))
		if err != nil {
			t.Log(err)
		}
	}(binaryName)

	//command := exec.Command("./ccrun", "run", "echo", "Hello Coding Challenges!")
	//command.Dir = "/home/emre/projects/ccrun"
	//
	//stdoutPipe, err := command.StdoutPipe()
	//assert.Nil(t, err)
	//
	//err = command.Start()
	//assert.Nil(t, err)
	//
	//stdoutMsg, err := io.ReadAll(stdoutPipe)
	//assert.Nil(t, err)
	//assert.Equal(t, "Hello Coding Challenges!\n", string(stdoutMsg))
	//
	//err = command.Wait()
	//assert.Nil(t, err)
}

func buildBinary(t *testing.T) string {
	binary := fmt.Sprintf("ccrun-%s", randomString(3))
	buildCmd := exec.Command("go", "build", "-o", binary)
	buildCmd.Dir = ".."
	var stdout, stderr strings.Builder
	buildCmd.Stdout = &stdout
	buildCmd.Stderr = &stderr
	err := buildCmd.Run()
	assert.Nil(t, err, "stdout:\n%s\nstderr:%s\n", stdout.String(), stderr.String())
	return binary
}

func randomString(length int) string {
	alphaNumericCharacters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)
	for i := range length {
		result[i] = alphaNumericCharacters[rand.IntN(len(alphaNumericCharacters))]
	}

	return string(result)
}
