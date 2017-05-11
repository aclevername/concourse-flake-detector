package concourse

import "fmt"

//go:generate counterfeiter -o fake/fake_client.go . ClientInterface
type ClientInterface interface {
	GetPipeline(string) (Pipeline, error)
	GetBuilds(Job) ([]Build, error)
	GetResources(Build) (Run, error)
}

//go:generate counterfeiter -o fake/fake_get.go . Getter
type Getter func(string) ([]byte, error)

type client struct {
	get     func(string) ([]byte, error)
	baseURL string
	url     string
	teamURL string
}

func NewClient(get Getter, baseURL, team string) *client {
	url := fmt.Sprintf("%s/api/v1", baseURL)

	teamURL := url
	if team != "" {
		teamURL = fmt.Sprintf("%s/teams/%s", teamURL, team)
	}

	return &client{
		get:     get,
		url:     url,
		baseURL: baseURL,
		teamURL: teamURL,
	}
}

// example team URL api/v1/teams/main/pipelines/main/jobs/fly/builds
