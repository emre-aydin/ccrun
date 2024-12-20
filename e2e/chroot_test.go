package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestChroot(t *testing.T) {
	binary := buildBinary(t)
	assert.NotNil(t, binary)
	t.Cleanup(func() {
		err := os.Remove(binary)
		if err != nil {
			t.Log(err)
		}
	})

	var stdout, stderr strings.Builder
	cmd := exec.Command(binary, "run", "/bin/busybox", "ls")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	assert.Nil(t, err)
	assert.Empty(t, stderr.String())
	assert.Contains(t, stdout.String(), "ALPINE_FS_ROOT")
}
