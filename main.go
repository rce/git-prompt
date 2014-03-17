package main

import (
	"bufio"
	"bytes"
	"fmt"
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

	// Read the first line containing branch name
	r := bufio.NewReader(b)
	var line string
	line, err = r.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	// Find out the branch name from the string
	branchRegexp := regexp.MustCompile("# On branch (.+)")
	branch := branchRegexp.FindStringSubmatch(line)[1]
	fmt.Println(branch)
}
