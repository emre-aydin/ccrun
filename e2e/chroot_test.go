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
	lsCurDirCmd := exec.Command(binary, "run", "ls")
	lsCurDirCmd.Stdout = &stdout
	lsCurDirCmd.Stderr = &stderr

	err := lsCurDirCmd.Run()
	assert.Nil(t, err)
	assert.Empty(t, err)
	curDirContent := stdout.String()

	lsParentDirCmd := exec.Command(binary, "run", "ls", "..")
	lsParentDirCmd.Stdout = &stdout
	lsParentDirCmd.Stderr = &stderr

	err = lsParentDirCmd.Run()
	assert.Nil(t, err)
	assert.Empty(t, err)

	parentDirContent := stdout.String()
	assert.Equal(t, curDirContent, parentDirContent)

	t.Log("cur dir content")
	t.Log(curDirContent)
	t.Log("parent dir content")
	t.Log(parentDirContent)
}
