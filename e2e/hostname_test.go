package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestHostname(t *testing.T) {
	binary := buildBinary(t)
	assert.NotNil(t, binary)
	t.Cleanup(func() {
		err := os.Remove(binary)
		if err != nil {
			t.Log(err)
		}
	})

	originalHostname, err := os.Hostname()
	assert.Nil(t, err)
	assert.NotEmpty(t, originalHostname)

	var stdout, stderr strings.Builder
	cmd := exec.Command(binary, "run", "hostname")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	assert.Nil(t, err)
	assert.NotEqual(t, originalHostname, strings.TrimSuffix(stdout.String(), "\n"))
	assert.Empty(t, stderr.String())
}
