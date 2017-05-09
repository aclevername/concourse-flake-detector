package historybuilder

import (
	"encoding/json"
	"fmt"
	"github.com/aclevername/concourse-flake-detector/api"
)

type Run struct {
	Status    string `json:"status"`
	Resources Resource
}

type Resource struct {
	Inputs []Input `json:"inputs"`
}

type Input struct {
	Name       string `json:"name"`
	Resource   string `json:"resource"`
	Type       string `json:"type"`
	Version    Ref    `json:"version"`
	PipelineID int    `json:"pipeline_id"`
}

type Ref struct {
	Ref string `json:"ref"`
}

type JSONBuild struct {
	Status string `json:"status"`
	APIURL string `json:"api_url"`
}

func GetJobHistory(client api.Client, job api.Job) ([]Run, error) {
	var err error
	response, _ := client.Get(fmt.Sprintf("api/v1%s/builds", job.URL))

	builds := make([]JSONBuild, 0)
	err = json.Unmarshal(response, &builds)
	if err != nil {
		panic(string(response))
	}
	var history []Run
	for _, build := range builds {
		response, _ := client.Get(build.APIURL + "/resources")

		inputs := Resource{}
		err = json.Unmarshal(response, &inputs)
		if err != nil {
			panic(string(response))
		}
		run := Run{
			Status:    build.Status,
			Resources: inputs,
		}

		history = append(history, run)
	}

	return history, nil
}
