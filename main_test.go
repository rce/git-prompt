package main

import "testing"

func TestParseStatus(t *testing.T) {
	b := `# On branch master
# Untracked files:
#   (use "git add <file>..." to include in what will be committed)
#
#	untracked_file
nothing added to commit but untracked files present (use "git add" to track)
`
	branch, err := readBranchFromStatus(b)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if branch != "master" {
		t.Errorf("Expected branch 'master', got '%v'", branch)
	}
}
