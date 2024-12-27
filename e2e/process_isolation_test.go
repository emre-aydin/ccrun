package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestProcessIsolation(t *testing.T) {
	binary := buildBinary(t)
	assert.NotNil(t, binary)
	t.Cleanup(func() {
		err := os.Remove(binary)
		if err != nil {
			t.Log(err)
		}
	})

	var stdout, stderr strings.Builder
	cmd := exec.Command(binary, "run", "/bin/busybox", "ps")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	assert.Nil(t, err)
	assert.Empty(t, stderr.String())
	assert.Contains(t, stdout.String(), "/bin/busybox ps")
}
