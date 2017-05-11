package mockconcourse

import (
	"fmt"

	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
)

type jobsMock struct {
	*mockhttp.Handler
}

func JobsForPipeline(pipeline, team string) *jobsMock {
	if team == "" {

		return &jobsMock{
			mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/api/v1/pipelines/%s/jobs", pipeline)),
		}
	} else {

		return &jobsMock{
			mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/api/v1/teams/%s/pipelines/%s/jobs", team, pipeline)),
		}
	}
}

func (j *jobsMock) RespondsWithJob(name, url string) *mockhttp.Handler {
	return j.RespondsOKWith(fmt.Sprintf(`[{"name":"%s","url":"%s"}]`, name, url))
}
