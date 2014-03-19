package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var tpl = flag.String("t", "{{.Branch}}", "Template for prompt string")

type gitInfo struct {
	Branch string // Name of the current branch or a commit hash if you are not on a specific branch
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
	b, err := runCommand("git", "symbolic-ref", "HEAD")
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(b.String(), "refs/heads/"), nil
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
	if err != nil {
		branch, err = currentHash()
		if err != nil {
			return nil, err
		}
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
