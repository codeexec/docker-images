package main

import (
	"flag"

	"github.com/kjk/u"
)

func main() {
	u.CdUpDir("codeeval-images")

	var (
		flgTests bool
	)

	flag.BoolVar(&flgTests, "tests", false, "run tests using docker image")

	if flgTests {
		runTests()
		return
	}

}
