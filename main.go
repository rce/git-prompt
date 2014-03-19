package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"text/template"
)

var tpl = flag.String("t", "{{.Branch}}", "Template for prompt string")

type gitStatus struct {
	Branch string
}

func main() {
	flag.Parse()

	b, err := runCommand("git", "status")
	if err != nil {
		os.Exit(1)
	}

	status, err := parseStatus(b)
	if err != nil {
		os.Exit(1)
	}

	prompt, err := makePrompt(status, *tpl)
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
	results := branchRegexp.FindStringSubmatch(line)
	if len(results) == 2 {
		// Found branch name
		status.Branch = results[1]
	} else {
		// Figure out commit name
		b, err := runCommand("git", "rev-parse", "--short", "HEAD")
		if err != nil {
			return nil, err
		}

		status.Branch = b.String()
	}

	return &status, nil
}

func makePrompt(status *gitStatus, tpl string) (string, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return "", nil
	}

	b := &bytes.Buffer{}
	err = t.Execute(b, status)
	return b.String(), nil
}
