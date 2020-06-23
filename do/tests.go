package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kjk/u"
)

const (
	testsDir       = "temp_tests"
	testScriptName = "testscript.sh"
)

// Test represents a single test case
type Test struct {
	fileName string
	run      []string
	cleanup  []string
	src      string
	output   string
	// unique n identifying a test, 0...N
	n   int
	raw string
}

func dirForTest(test *Test) string {
	return filepath.Join(testsDir, fmt.Sprintf("%02d", test.n))
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
		test.fileName = strings.TrimSpace(s)
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
	var res Test
	res.raw = s
	parts := strings.Split(s, "\n--\n")
	u.PanicIf(len(parts) != 2, "len(parts) == %d in:\n..%s..\n", len(parts), s)
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

func validateTest(test *Test) {
	panicIf(test.raw == "")
	panicIf(test.fileName == "", "test:\n%s\n", test.raw)
	panicIf(test.src == "", "test:\n%s\n", test.raw)
	panicIf(test.output == "", "test:\n%s\n", test.raw)
	panicIf(len(test.run) == 0, "test:\n%s\n", test.raw)
}

func buildTestScript(test *Test) string {
	s := `#/bin/bash
set -eux

`

	for idx, run := range test.run {
		run = strings.Replace(run, "$file", test.fileName, -1)
		isLast := idx == len(test.run)-1
		// we assume last command is the execution, so redirect the output to a file
		// for comparison with expected
		if isLast {
			run += " >output.txt"
		}
		s += run + "\n"
	}

	s += `
diff --ignore-trailing-space --unified output.txt exp_output.txt
`
	// TODO: do I need cleanup?
	return s
}

func writeOutTest(test *Test, dir string) {
	u.CreateDirMust(dir)
	path := filepath.Join(dir, "exp_output.txt")
	u.WriteFileMust(path, []byte(test.output))
	path = filepath.Join(dir, test.fileName)
	u.WriteFileMust(path, []byte(test.src))
	s := buildTestScript(test)
	path = filepath.Join(dir, testScriptName)
	u.WriteFileMust(path, []byte(s))
}

func writeOutTests(tests []*Test) {
	for _, test := range tests {
		testDir := dirForTest(test)
		writeOutTest(test, testDir)
	}
}

func deleteTests(dir string) {
	must(os.RemoveAll(dir))
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
		validateTest(test)
		test.n = len(res)
		res = append(res, test)
	}
	return res
}

func runTests() {
	path := filepath.Join("do", "tests.txt")
	tests := loadTests(path)
	fmt.Printf("%d tests\n", len(tests))

	deleteTests(testsDir)
	writeOutTests(tests)

	for _, test := range tests {
		fmt.Printf("test:\n%s\n---\n", test.src)
		dir, err := filepath.Abs(dirForTest(test))
		must(err)
		v := fmt.Sprintf("%s:/test", dir)
		// testScriptName
		cmd := exec.Command("docker", "run", "--rm", "-v", v, "-w=/test", "eval-multi-base:latest", "/bin/bash -c ./testscript.sh")
		u.RunCmdLoggedMust(cmd)
	}
}
