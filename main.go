package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"text/template"
)

var tpl = flag.String("t", "{{.Branch}}", "Template for prompt string")

var ErrNotOnBranch = errors.New("you are not currently on a commit tagged as a branch")

type gitInfo struct {
	Branch string
}

func main() {
	flag.Parse()

	info, err := getInfo()
	if err != nil {
		os.Exit(1)
	}

	prompt, err := makePrompt(info, *tpl)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(prompt)
}

func runCommand(command string, args ...string) (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	cmd := exec.Command(command, args...)
	cmd.Stdout = b
	err := cmd.Run()
	return b, err
}

func currentBranch() (string, error) {
	b, err := runCommand("git", "status")
	if err != nil {
		os.Exit(1)
	}

	return readBranchFromStatus(b.String())
}

func readBranchFromStatus(status string) (string, error) {
	branchRegexp := regexp.MustCompile("# On branch (.+)")
	results := branchRegexp.FindStringSubmatch(status)
	if len(results) != 2 {
		return "", ErrNotOnBranch
	}
	return results[1], nil
}

func currentHash() (string, error) {
	b, err := runCommand("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func getInfo() (*gitInfo, error) {
	var info gitInfo

	// Find branch or current commit hash
	branch, err := currentBranch()
	if err == ErrNotOnBranch {
		branch, err = currentHash()
	}
	if err != nil {
		return nil, err
	}

	info.Branch = branch

	return &info, nil
}

func makePrompt(status *gitInfo, tpl string) (string, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return "", nil
	}

	b := &bytes.Buffer{}
	err = t.Execute(b, status)
	return b.String(), nil
}
