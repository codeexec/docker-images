package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kjk/u"
)

// Test represents a single test case
type Test struct {
	file    string
	run     []string
	cleanup []string
	src     string
	output  string
}

func eatPrefix(s string, prefix string) (string, bool) {
	res := strings.TrimPrefix(s, prefix)
	return res, res != s
}

/*
:file main.swift
:run swiftc $file -o main
:run ./main
:cleanup rm ./main
print("hello from swift")
--
hello from swift
*/
func parseDirective(s string, test *Test) {
	if s, ok := eatPrefix(s, ":file "); ok {
		test.file = strings.TrimSpace(s)
		return
	}
	if s, ok := eatPrefix(s, ":run "); ok {
		test.run = append(test.run, strings.TrimSpace(s))
		return
	}
	if s, ok := eatPrefix(s, ":cleanup "); ok {
		test.cleanup = append(test.cleanup, strings.TrimSpace(s))
		return
	}
	panicIf(true, "unknown directive in line:\n%s\n\n", s)
}

func parseTest(s string) *Test {
	parts := strings.Split(s, "\n--\n")
	u.PanicIf(len(parts) != 2, "len(parts) == %d in:\n..%s..\n", len(parts), s)
	var res Test
	res.output = parts[1]
	lines := strings.Split(parts[0], "\n")
	panicIf(len(lines) < 3)
	for len(lines) > 0 {
		s := lines[0]
		if s[0] != ':' {
			break
		}
		parseDirective(s, &res)
		lines = lines[1:]
	}
	panicIf(len(lines) == 0)
	res.src = strings.Join(lines, "\n")
	return &res
}

func loadTests(path string) []*Test {
	d := u.ReadFileMust(path)
	d = u.NormalizeNewlines(d)
	s := string(d)
	tests := strings.Split(s, "\n---\n")
	panicIf(len(tests) < 2)
	fmt.Printf("len(tests): %d\n", len(tests))
	var res []*Test
	for n, testStr := range tests {
		// skip empty string if at the end
		s := strings.TrimSpace(testStr)
		if len(s) == 0 {
			if n == len(tests)-1 {
				continue
			}
			panicIf(true, "empty test at post %d out of %d\n", n, len(tests))
		}
		test := parseTest(testStr)
		res = append(res, test)
	}
	return res
}

func runTest(test *Test) {
	panicIf(len(test.run) == 0, "test: %#v", test)
	fmt.Printf("test: %s\n", test.src)
}

func runTests() {
	path := filepath.Join("do", "tests.txt")
	tests := loadTests(path)
	fmt.Printf("%d tests\n", len(tests))
	for _, test := range tests {
		runTest(test)
	}
}
