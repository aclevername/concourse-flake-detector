package mockconcourse

import (
	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
	"fmt"
)

type jobsMock struct{
	*mockhttp.Handler
}

func JobsForPipeline(pipeline string) *jobsMock {
	return &jobsMock{
		mockhttp.NewMockedHttpRequest("GET", fmt.Sprintf("/api/v1/pipelines/%s/jobs", pipeline)),
	}
}

func (j *jobsMock) RespondsWithJob(name, url string) *mockhttp.Handler {
	job := struct {
		Name string
		url string
	}{Name: name, url: url}

	var listJobs []struct {
		Name string
		url string
	}

	listJobs = append(listJobs, job)
	return j.RespondsOKWithJSON(listJobs)
}


