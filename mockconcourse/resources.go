package mockconcourse

import (
	"fmt"
	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
)

type resourcesMock struct {
	*mockhttp.Handler
}

func ResourcesForBuild(build string) *resourcesMock {
	return &resourcesMock{
		mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("%s/resources", build)),
	}
}

func (j *resourcesMock) RespondsWith(builds string) *mockhttp.Handler {

	return j.RespondsOKWith(builds)
}
