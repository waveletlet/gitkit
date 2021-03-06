package gitkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseGitCommand(t *testing.T) {
	examples := map[string]GitCommand{
		"git-upload-pack 'hello.git'":    GitCommand{"git-upload-pack", "hello.git", "git-upload-pack 'hello.git'"},
		"git upload-pack 'hello.git'":    GitCommand{"git upload-pack", "hello.git", "git upload-pack 'hello.git'"},
		"git-receive-pack 'hello.git'":   GitCommand{"git-receive-pack", "hello.git", "git-receive-pack 'hello.git'"},
		"git receive-pack 'hello.git'":   GitCommand{"git receive-pack", "hello.git", "git receive-pack 'hello.git'"},
		"git-upload-archive 'hello.git'": GitCommand{"git-upload-archive", "hello.git", "git-upload-archive 'hello.git'"},
		"git upload-archive 'hello.git'": GitCommand{"git upload-archive", "hello.git", "git upload-archive 'hello.git'"},
	}

	for s, expected := range examples {
		cmd, err := ParseGitCommand(s)

		assert.NoError(t, err)
		assert.Equal(t, expected.Command, cmd.Command)
		assert.Equal(t, expected.Repo, cmd.Repo)
	}

	cmd, err := ParseGitCommand("git do-stuff")
	assert.Error(t, err)
	assert.Nil(t, cmd)
}
