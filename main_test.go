package main

import (
	"bytes"
	"fmt"
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
		"error\t\tpath/to/fileb.go:40:\t\tFoo bar",
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
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "error\tpath/to/filea.go:10:4:\tsss ssssd sd")
	fmt.Fprintln(w, "warning\tpath/to/fileb.go:4:4:\tfdsqfds fdsq")
	fmt.Fprintln(w, "error\tpath/to/fileb.go:10:\tFoo bar")

	w.Flush()

	prob := generateProblemMatcher()

	pattern := prob.Matchers[0].Pattern[0]

	createReplacement(pattern)

	exp := regexp.MustCompile(pattern.Regexp)

	lines := strings.Split(buf.String(), "\n")

	for _, line := range lines {
		fmt.Println()
		fmt.Println(line)
		fmt.Println(exp.MatchString(line))

		fmt.Println(exp.ReplaceAllString(line, createReplacement(pattern)))
	}

}
