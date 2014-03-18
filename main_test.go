package main

import (
	"bytes"
	"testing"
)

func TestParseStatus(t *testing.T) {
	b := bytes.NewBufferString(`# On branch master
# Untracked files:
#   (use "git add <file>..." to include in what will be committed)
#
#	untracked_file
nothing added to commit but untracked files present (use "git add" to track)
`)
	status, err := parseStatus(b)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if status.Branch != "master" {
		t.Errorf("Expected branch 'master', got '%v'", status.Branch)
	}
}
