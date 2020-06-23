package main

import (
	"strings"

	"github.com/kjk/u"
)

type test struct {
	run     []string
	cleanup []string
	test    string
	output  string
}

func loadTests(path string) []test {
	d := u.ReadFileMust(path)
	d = u.NormalizeNewlines(d)
	s := string(d)
	tests := strings.Split(s, "---\n")

	return nil
}

func runTests() {
	tests := loadTests("tests.txt")
	for _, test := range tests {

	}
}
