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

	// Note: the file with the problem matcher definition should not be removed.
	//
	// Error: Unable to process command '::add-matcher::/tmp/golangci-lint-action-1296491320-problem-matchers.json' successfully.
	// Error: Could not find file '/tmp/golangci-lint-action-1296491320-problem-matchers.json'.

	fmt.Printf("::debug::problem matcher definition file: %s\n", filename)

	fmt.Println("::group::Linting Issues")

	fmt.Printf("::add-matcher::%s\n", filename)

	fmt.Println("path/to/filea.go:10:4:\t[error]\tsss ssssd sd")
	fmt.Println("path/to/fileb.go:1:4:\t[warning]\tfdsqfds fdsq")
	fmt.Println("path/to/fileb.go:40:4:\t[error]\tFoo bar")

	fmt.Println("::endgroup::")

	fmt.Println("::remove-matcher owner=golangci-lint-action::")

	return nil
}

func storeProblemMatcher() (string, error) {
	prob := ProblemMatcher{
		Matchers: []Matcher{
			{
				Owner:    "golangci-lint-action",
				Severity: "error",
				Pattern: []Pattern{
					{
						Regexp:   `^(.+):(\d+):(\d+):\t\[(.+)\]\t(.+)$`,
						File:     1,
						Line:     2,
						Column:   3,
						Severity: 4,
						Message:  5,
					},
				},
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
	Matchers []Matcher `json:"problemMatcher,omitempty"`
}

type Matcher struct {
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
