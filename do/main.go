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

func main() {
	var (
		flgRunTests    bool
		flgBuildGcr    bool
		flgDockerBuild bool
	)

	flag.BoolVar(&flgRunTests, "run-tests", false, "run tests inside docker image")
	flag.BoolVar(&flgBuildGcr, "build-gcr", false, "submit to GCR for build")
	flag.BoolVar(&flgDockerBuild, "docker-build", false, "build the image locally with docker")
	flag.Parse()

	timeStart := time.Now()

	if flgRunTests {
		// inside docker our parent dir /home/runner
		u.CdUpDir("runner")
		runTests()
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	u.CdUpDir("codeeval-images")

	if flgBuildGcr {
		buildGcr()
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	if flgDockerBuild {
		dockerBuild()
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	flag.Usage()
}
