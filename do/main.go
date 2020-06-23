package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/kjk/u"
)

var panicIf = u.PanicIf
var panicIfErr = u.PanicIfErr
var must = u.Must

func dockerBuild() {
	cmd := exec.Command("docker", "build", "-f", "Dockerfile", "-t", "eval-multi-base:latest", ".")
	u.RunCmdLoggedMust(cmd)
}

func dockerRun() {
	cmd := exec.Command("docker", "run", "--rm", "-it", "eval-multi-base:latest", "/bin/bash")
	u.RunCmdMust(cmd)
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

	timeStart := time.Now()

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
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	if flgDockerRun {
		dockerRun()
		return
	}

	flag.Usage()
}
