package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	prob := generateProblemMatcher()

	pattern := prob.Matchers[0].Pattern[0]

	createReplacement(pattern)

	exp := regexp.MustCompile(pattern.Regexp)

	lines := []string{
		"path/to/filea.go\t10:4:\t[error]\tsss ssssd sd",
		"path/to/fileb.go\t1:4:\t[warning]\tfdsqfds fdsq",
		"path/to/fileb.go\t40:4:\t[error]\tFoo bar",
	}

	for _, line := range lines {
		fmt.Println()
		fmt.Println(exp.ReplaceAllString(line, createReplacement(pattern)))
	}
}

func createReplacement(pattern Pattern) string {
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
