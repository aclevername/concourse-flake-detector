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
	Name string `json:"name"`
	URL  string `json:"api_url"`
}

func GetPipeline(url, name string, client Client) (Pipeline, error) {
	response, err := client.Get(fmt.Sprintf("api/v1/pipelines/%s/jobs", name))
	if err != nil {
		return Pipeline{}, err
	}

	fmt.Println("++++++++++" + string(response))
	jobs := make([]Job, 0)
	err = json.Unmarshal(response, &jobs)

	return Pipeline{jobs: jobs}, err
}

func (p *Pipeline) Jobs() []Job {
	return p.jobs
}
