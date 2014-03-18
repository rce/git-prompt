package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
)

type gitStatus struct {
	Branch string
}

func main() {
	r, err := runGitStatus()
	if err != nil {
		os.Exit(1)
	}

	status, err := parseStatus(r)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(makePrompt(status))
}

func runGitStatus() (io.Reader, error) {
	b := &bytes.Buffer{}
	cmd := exec.Command("git", "status")
	cmd.Stdout = b
	err := cmd.Run()
	return b, err
}

func parseStatus(src io.Reader) (*gitStatus, error) {
	r := bufio.NewReader(src)
	var status gitStatus

	// Read the first line containing branch name
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Find out the branch name from the string
	branchRegexp := regexp.MustCompile("# On branch (.+)")
	status.Branch = branchRegexp.FindStringSubmatch(line)[1]

	return &status, nil
}

func makePrompt(status *gitStatus) string {
	return status.Branch
}
