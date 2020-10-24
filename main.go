package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/kjk/u"
)

var panicIf = u.PanicIf
var panicIfErr = u.PanicIfErr
var must = u.Must

func dockerBuild(noCache bool) {
	var cmd *exec.Cmd
	if noCache {
		cmd = exec.Command("docker", "build", "--no-cache", "-f", "Dockerfile-20.04", "-t", "eval-multi-base-20_04:latest", ".")
	} else {
		cmd = exec.Command("docker", "build", "-f", "Dockerfile-20.04", "-t", "eval-multi-base-20_04:latest", ".")
	}
	u.RunCmdLoggedMust(cmd)
}

const (
	// https://console.cloud.google.com/home/dashboard?folder=&organizationId=&project=cloudeval-255302
	gcpProject = "cloudeval-255302"
)

func buildGcr() {
	// https://cloud.google.com/sdk/gcloud/reference/builds/submit
	{
		configPath := filepath.Join("cloudbuild.yml")
		cmd := exec.Command("gcloud", "builds", "submit", "--project", gcpProject, "--config", configPath, "--timeout=3600s")
		u.RunCmdLoggedMust(cmd)
	}
}

func main() {
	var (
		flgRunTests    bool
		flgBuildGcr    bool
		flgDockerBuild bool
		flgNoCache     bool
	)

	flag.BoolVar(&flgRunTests, "run-tests", false, "run tests inside docker image")
	flag.BoolVar(&flgBuildGcr, "build-gcr", false, "submit to GCR for build")
	flag.BoolVar(&flgDockerBuild, "build-docker", false, "build the image locally with docker")
	flag.BoolVar(&flgNoCache, "no-cache", false, "if true, disable docker build cache")
	flag.Parse()

	timeStart := time.Now()

	if flgRunTests {
		// inside docker our parent dir /home/runner
		u.CdUpDir("runner")
		runTests()
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	u.CdUpDir("docker-images")

	if flgBuildGcr {
		buildGcr()
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	if flgDockerBuild {
		dockerBuild(flgNoCache)
		fmt.Printf("Took %s\n", time.Since(timeStart))
		return
	}

	flag.Usage()
}

func runTests() {
	// impelement me or remove me
}
