package main

import (
	"bytes"
	"testing"
)

func TestSingleUntrackedFileOnBranchMaster(t *testing.T) {
	const singleUntrackedFile = `# On branch master
# Untracked files:
#   (use "git add <file>..." to include in what will be committed)
#
#	untracked_file
nothing added to commit but untracked files present (use "git add" to track)
`
	s, err := makePrompt(bytes.NewBufferString(singleUntrackedFile))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if s != "master" {
		t.Errorf("Expected prompt string 'master', got '%v'", s)
	}
}
