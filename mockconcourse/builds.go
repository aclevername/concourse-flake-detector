package mockconcourse

import (
	"fmt"

	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
)

type buildMock struct {
	*mockhttp.Handler
}

func BuildsForJob(job string) *buildMock {
	return &buildMock{
		mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/api/v1%s/builds", job)),
	}
}

func (j *buildMock) RespondsWithBuilds(builds string) *mockhttp.Handler {

	return j.RespondsOKWith(builds)
}
