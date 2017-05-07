package job

import "github.com/aclevername/concourse-flake-detector/build"

type Job struct {
	Name   string
	URL    string
	builds []build.Build
}
