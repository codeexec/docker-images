package main

import (
	"flag"

	"github.com/kjk/u"
)

var panicIf = u.PanicIf
var panicIfErr = u.PanicIfErr

func dockerBuild() {
	u.RunCmdMust("docker", "build", "-f", "Dockerfile", "-t", "eval-multi-base:latest", ".")
}

func dockerRun() {
	u.RunCmdMust("docker", "run", "--rm", "-it", "eval-multi-base:latest", "/bin/bash")
}

func main() {
	u.CdUpDir("codeeval-images")

	var (
		flgRunTests    bool
		flgBuildGcr    bool
		flgDockerBuild bool
		flgDockerRun   bool
	)

	flag.BoolVar(&flgRunTests, "run-tests", false, "run tests using docker image")
	flag.BoolVar(&flgBuildGcr, "build-gcr", false, "submit to GCR for build")
	flag.BoolVar(&flgDockerBuild, "docker-build", false, "build the image locally with docker")
	flag.BoolVar(&flgDockerRun, "docker-run", false, "run the image locally with docker")
	flag.Parse()

	if flgRunTests {
		runTests()
		return
	}

	if flgBuildGcr {
		buildGcr()
		return
	}

	if flgDockerBuild {
		dockerBuild()
		return
	}

	if flgDockerRun {
		dockerRun()
		return
	}

	flag.Usage()
}
