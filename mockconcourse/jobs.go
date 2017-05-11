package mockconcourse

import (
	"fmt"

	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
)

type jobsMock struct {
	*mockhttp.Handler
}

func JobsForPipeline(pipeline string) *jobsMock {
	return &jobsMock{
		mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/api/v1/pipelines/%s/jobs", pipeline)),
	}
}

func (j *jobsMock) RespondsWithJob(name, url string) *mockhttp.Handler {
	return j.RespondsOKWith(fmt.Sprintf(`[{"name":"%s","api_url":"%s"}]`, name, url))
}
