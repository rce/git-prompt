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

func main() {
	// Buffer for git status output
	b := &bytes.Buffer{}

	// Run command
	cmd := exec.Command("git", "status")
	cmd.Stdout = b
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
	prompt, err := makePrompt(b)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(prompt)
}

func makePrompt(status io.Reader) (string, error) {
	// Read the first line containing branch name
	r := bufio.NewReader(status)
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Find out the branch name from the string
	branchRegexp := regexp.MustCompile("# On branch (.+)")
	branch := branchRegexp.FindStringSubmatch(line)[1]

	return branch, nil
}
