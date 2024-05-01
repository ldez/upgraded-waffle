package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"text/tabwriter"
)

func TestName(t *testing.T) {
	prob := generateProblemMatcher()

	pattern := prob.Matchers[0].Pattern[0]

	createReplacement(pattern)

	exp := regexp.MustCompile(pattern.Regexp)

	lines := []string{
		"error\tpath/to/filea.go:10:4:\tsome issue (sample-linter)",
		"warning\tpath/to/fileb.go:1:4:\tsome issue (sample-linter)",
		"error\tpath/to/fileb.go:40:\tFoo bar",
	}

	for _, line := range lines {
		fmt.Println()
		fmt.Println(exp.ReplaceAllString(line, createReplacement(pattern)))
	}
}

func createReplacement(pattern GitHubPattern) string {
	var repl []string

	if pattern.File > 0 {
		repl = append(repl, fmt.Sprintf("File: $%d", pattern.File))
	}

	if pattern.FromPath > 0 {
		repl = append(repl, fmt.Sprintf("FromPath: $%d", pattern.FromPath))
	}

	if pattern.Line > 0 {
		repl = append(repl, fmt.Sprintf("Line: $%d", pattern.Line))
	}

	if pattern.Column > 0 {
		repl = append(repl, fmt.Sprintf("Column: $%d", pattern.Column))
	}

	if pattern.Severity > 0 {
		repl = append(repl, fmt.Sprintf("Severity: $%d", pattern.Severity))
	}

	if pattern.Code > 0 {
		repl = append(repl, fmt.Sprintf("Code: $%d", pattern.Code))
	}

	if pattern.Message > 0 {
		repl = append(repl, fmt.Sprintf("Message: $%d", pattern.Message))
	}

	if pattern.Loop {
		repl = append(repl, fmt.Sprintf("Loop: $%v", pattern.Loop))
	}

	return strings.Join(repl, "\n")
}

func TestWriter(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "error\tpath/to/filea.go:10:4:\tsss ssssd sd")
	fmt.Fprintln(w, "warning\tpath/to/fileb.go:1:4:\tfdsqfds fdsq")
	fmt.Fprintln(w, "error\tpath/to/fileb.go:40:\tFoo bar")

	w.Flush()

}
