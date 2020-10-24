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

func dockerBuild() {
	cmd := exec.Command("docker", "build", "-f", "Dockerfile", "-t", "eval-multi-base:latest", ".")
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

	u.CdUpDir("docker-images")

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

func runTests() {
	// impelement me or remove me
}
