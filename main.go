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
	// A sleep can mitigate this problem but this will be flaky.
	//
	// Error: Unable to process command '::add-matcher::/tmp/golangci-lint-action-1296491320-problem-matchers.json' successfully.
	// Error: Could not find file '/tmp/golangci-lint-action-1296491320-problem-matchers.json'.

	fmt.Printf("::debug::problem matcher definition file: %s\n", filename)

	fmt.Println("::group::Linting Issues")

	fmt.Printf("::add-matcher::%s\n", filename)

	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	//
	// fmt.Fprintln(w, "error\tpath/to/filea.go:10:4:\tsss ssssd sd")
	// fmt.Fprintln(w, "warning\tpath/to/fileb.go:1:4:\tfdsqfds fdsq")
	// fmt.Fprintln(w, "error\tpath/to/fileb.go:40:\tFoo bar")
	//
	// w.Flush()

	fmt.Println("error\tpath/to/filea.go:10:4:\tsss ssssd sd")
	fmt.Println("warning\tpath/to/fileb.go:1:4:\tfdsqfds fdsq")
	fmt.Println("error\t\tpath/to/fileb.go:40:\tFoo bar")

	fmt.Println("::endgroup::")

	fmt.Println("::remove-matcher owner=golangci-lint-action::")

	return nil
}

func storeProblemMatcher() (string, error) {
	file, err := os.CreateTemp("", "golangci-lint-action-*-problem-matchers.json")
	if err != nil {
		return "", err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(generateProblemMatcher())
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func generateProblemMatcher() GitHubProblemMatchers {
	return GitHubProblemMatchers{
		Matchers: []GitHubMatcher{
			{
				Owner:    "golangci-lint-action",
				Severity: "error",
				Pattern: []GitHubPattern{
					{
						Regexp:   `^([^\t]+)\t+([^\t]+):(\d+):(?:(\d+):)?\t+(.+)$`,
						File:     2,
						Line:     3,
						Column:   4,
						Severity: 1,
						Message:  5,
					},
				},
			},
		},
	}
}

// GitHubProblemMatchers defines the root of problem matchers.
// https://github.com/actions/toolkit/blob/main/docs/problem-matchers.md
type GitHubProblemMatchers struct {
	Matchers []GitHubMatcher `json:"problemMatcher,omitempty"`
}

// GitHubMatcher defines a problem matcher.
type GitHubMatcher struct {
	// Owner an ID field that can be used to remove or replace the problem matcher.
	// **required**
	Owner string `json:"owner,omitempty"`
	// Severity indicates the default severity, either 'warning' or 'error' case-insensitive.
	// Defaults to 'error'.
	Severity string          `json:"severity,omitempty"`
	Pattern  []GitHubPattern `json:"pattern,omitempty"`
}

// GitHubPattern defines a pattern for a problem matcher.
type GitHubPattern struct {
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
