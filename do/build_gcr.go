package main

import (
	"os/exec"
	"path/filepath"

	"github.com/kjk/u"
)

const (
	// https://console.cloud.google.com/home/dashboard?folder=&organizationId=&project=cloudeval-255302
	gcpProject = "cloudeval-255302"
)

func buildGcr() {

	// https://cloud.google.com/sdk/gcloud/reference/builds/submit
	{
		configPath := filepath.Join("cloudbuild.yml")
		cmd := exec.Command("gcloud", "builds", "submit", "--project", gcpProject, "--config", configPath)
		u.RunCmdLoggedMust(cmd)
	}
}
