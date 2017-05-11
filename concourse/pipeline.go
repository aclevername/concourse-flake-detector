package concourse

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
	URL  string `json:"url"`
}

func (c *client) GetPipeline(name string) (Pipeline, error) {
	response, err := c.get(fmt.Sprintf("%s/pipelines/%s/jobs", c.teamURL, name))
	fmt.Println(string(response))
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
