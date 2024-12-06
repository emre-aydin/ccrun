package e2e_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommandLine(t *testing.T) {
	binary := buildBinary(t)
	assert.NotNil(t, binary)
	t.Cleanup(func() {
		err := os.Remove(fmt.Sprintf("../%s", binary))
		if err != nil {
			t.Log(err)
		}
	})

	tests := map[string]struct {
		Args             []string
		ExpectedStdout   string
		ExpectedStderr   string
		ExpectedExitCode int
	}{
		"execute run command": {
			Args:             []string{"run", "echo", "Hello coding challenges!"},
			ExpectedStdout:   "Hello coding challenges!\n",
			ExpectedStderr:   "",
			ExpectedExitCode: 0,
		},
		"execute run with no command": {
			Args:           []string{"run"},
			ExpectedStdout: "",
			ExpectedStderr: "no command to run\n",
			ExpectedExitCode: 1,
		},
		"execute invalid command": {
			Args:           []string{"exec", "echo", "Hello coding challenges!"},
			ExpectedStdout: "",
			ExpectedStderr: "invalid command: exec\n",
			ExpectedExitCode: 1,
		},
		"no command specified": {
			Args:           []string{},
			ExpectedStdout: "",
			ExpectedStderr: "no command specified: valid commands: [run]\n",
			ExpectedExitCode: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			stdout, stderr, err := executeCmd(t, binary, test.Args...)
			assert.Equal(t, test.ExpectedStdout, stdout)
			assert.Equal(t, test.ExpectedStderr, stderr)
			if test.ExpectedExitCode != 0 {
				var exitError *exec.ExitError
				assert.ErrorAs(t, err, &exitError)
				assert.Equal(t, test.ExpectedExitCode, exitError.ExitCode())
			} else {
				assert.Nil(t, err)
			}
		})
	}
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

func executeCmd(t *testing.T, binary string, args ...string) (string, string, error) {
	dir, err := os.Getwd()
	assert.Nil(t, err)
	command := exec.Command(filepath.Join(dir, "..", binary), args...)
	var stdout, stderr strings.Builder
	command.Stdout = &stdout
	command.Stderr = &stderr
	err = command.Run()
	return stdout.String(), stderr.String(), err
}

func randomString(length int) string {
	alphaNumericCharacters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)
	for i := range length {
		result[i] = alphaNumericCharacters[rand.IntN(len(alphaNumericCharacters))]
	}

	return string(result)
}
