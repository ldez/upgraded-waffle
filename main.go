package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	filename, err := storeProblemMatcher()
	if err != nil {
		return err
	}

	// defer os.RemoveAll(filename)

	fmt.Println(filename)

	fmt.Printf("::add-matcher::%s\n", filename)
	fmt.Println("::group::My title")

	fmt.Println("file=path/to/filea.go, line=10, col=4, linter=XXX, severity=error, message=sss ssssd sd")
	fmt.Println("file=path/to/fileb.go, line=1, col=4, linter=YYY, severity=warning, message=fdsqfds fdsq")

	fmt.Println("::endgroup::")

	// fmt.Println("::remove-matcher owner=golangci-lint-action::")

	return nil
}

func storeProblemMatcher() (string, error) {
	prob := ProblemMatcher{
		Owner:    "golangci-lint-action",
		Severity: "error",
		Pattern: []Pattern{
			{
				Regexp:   `^file=(.+), line=(\d+), col=(\d+), linter=(.+), severity=(.+), message=(.+)$`,
				File:     1,
				FromPath: 0, // ?
				Line:     2,
				Column:   3,
				Severity: 5,
				Code:     4,
				Message:  6,
			},
		},
	}

	file, err := os.CreateTemp("", "golangci-lint-action-*-problem-matchers.json")
	if err != nil {
		return "", err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(prob)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

type ProblemMatcher struct {
	// Owner an ID field that can be used to remove or replace the problem matcher.
	// **required**
	Owner string `json:"owner,omitempty"`
	// Severity indicates the default severity, either 'warning' or 'error' case-insensitive.
	// Defaults to 'error'.
	Severity string    `json:"severity,omitempty"`
	Pattern  []Pattern `json:"pattern,omitempty"`
}

type Pattern struct {
	// Regexp the regexp pattern that provides the groups to match against.
	// **required**
	Regexp string `json:"regexp,omitempty"`
	// File a group number containing the file name.
	File int `json:"file,omitempty"`
	// FromPath a group number containing a filepath used to root the file (e.g. a project file).
	FromPath int `json:"fromPath,omitempty"`
	// Line a group number containing the line number.
	Line int `json:"line,omitempty"`
	// Column a group number containing the column information.
	Column int `json:"column,omitempty"`
	// Severity a group number containing either 'warning' or 'error' case-insensitive.
	// Defaults to `error`.
	Severity int `json:"severity,omitempty"`
	// Code a group number containing the error code.
	Code int `json:"code,omitempty"`
	// Message a group number containing the error message.
	// **required** at least one pattern must set the message.
	Message int `json:"message,omitempty"`
	// Loop whether to loop until a match is not found,
	// only valid on the last pattern of a multi-pattern matcher.
	Loop bool `json:"loop,omitempty"`
}
