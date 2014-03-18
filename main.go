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

	r, err := runGitStatus()
	if err != nil {
		os.Exit(1)
	}

	status, err := parseStatus(r)
	if err != nil {
		os.Exit(1)
	}

	prompt, err := makePrompt(status, *tpl)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(prompt)
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

func makePrompt(status *gitStatus, tpl string) (string, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return "", nil
	}

	b := &bytes.Buffer{}
	err = t.Execute(b, status)
	return b.String(), nil
}
