package api

import (
	"encoding/json"
	"fmt"
)

type Pipeline struct {
	name string
	jobs []Job
}

type Job struct {
	Name string
	URL  string
}

func NewPipeline(url, name string, client Client) (Pipeline, error) {
	response, err := client.Get(fmt.Sprintf("%s/api/v1/pipelines/%s/jobs", url, name))
	if err != nil {
		return Pipeline{}, err
	}

	jobs := make([]Job, 0)
	err = json.Unmarshal(response, &jobs)

	return Pipeline{jobs: jobs}, err
}

func (p *Pipeline) Jobs() []Job {
	return p.jobs
}
