package pipeline

import (
	"encoding/json"
	"fmt"

)

type Pipeline struct {
	name string
	jobs []Job
}

type Job struct {
	Name   string
	URL    string
}


type HTTPClient interface {
	Get(string) ([]byte, error)
}

func New(url, name string, client HTTPClient) Pipeline {
	response, _ := client.Get(fmt.Sprintf("%s/api/v1/pipelines/%s/jobs", url, name))
	// if err != nil {
	// 	panic("failed to form request")
	// }

	jobs := make([]Job, 0)
	_ = json.Unmarshal(response, &jobs)

	return Pipeline{ jobs : jobs}
}

func (p *Pipeline) Jobs() []Job {
	return p.jobs
}