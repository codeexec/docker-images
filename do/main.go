package main

import (
	"flag"

	"github.com/kjk/u"
)

var panicIf = u.PanicIf
var panicIfErr = u.PanicIfErr

func main() {
	u.CdUpDir("codeeval-images")

	var (
		flgRunTests bool
		flgBuildGcr bool
	)

	flag.BoolVar(&flgRunTests, "run-tests", false, "run tests using docker image")
	flag.BoolVar(&flgBuildGcr, "build-gcr", false, "submit to GCR for build")
	flag.Parse()

	if flgRunTests {
		runTests()
		return
	}

	if flgBuildGcr {
		buildGcr()
		return
	}

	flag.Usage()
}
