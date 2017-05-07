package pipeline

import (
	"encoding/json"
	"fmt"

	"github.com/aclevername/concourse-flake-detector/job"
)

type Pipeline struct {
	name string
	jobs []job.Job
}

type HTTPClient interface {
	Get(string) ([]byte, error)
}

func New(url, name string, client HTTPClient) Pipeline {
	response, _ := client.Get(fmt.Sprintf("%s/api/v1/pipelines/%s/jobs", url, name))
	// if err != nil {
	// 	panic("failed to form request")
	// }

	jobs := make([]job.Job, 0)
	_ = json.Unmarshal(response, &jobs)

	return Pipeline{ jobs : jobs}
}

func (p *Pipeline) Jobs() []job.Job {
	return p.jobs

}
