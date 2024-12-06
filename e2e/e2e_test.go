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
		Msg string
	}{
		"a message": {
			Msg: "Hello coding challenges!",
		},
		"another message": {
			Msg: "Hello World!",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stdout, stderr := executeBinary(t, binary, test.Msg)
			assert.Empty(t, stderr)
			assert.Equal(t, fmt.Sprintf("%s\n", test.Msg), stdout)
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

func executeBinary(t *testing.T, binary, msg string) (string, string) {
	dir, err := os.Getwd()
	assert.Nil(t, err)
	command := exec.Command(filepath.Join(dir, "..", binary), "run", "echo", msg)
	var stdout, stderr strings.Builder
	command.Stdout = &stdout
	command.Stderr = &stderr
	err = command.Run()
	assert.Nil(t, err)
	return stdout.String(), stderr.String()
}

func randomString(length int) string {
	alphaNumericCharacters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)
	for i := range length {
		result[i] = alphaNumericCharacters[rand.IntN(len(alphaNumericCharacters))]
	}

	return string(result)
}
