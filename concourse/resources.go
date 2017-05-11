package concourse

import (
	"encoding/json"
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

func (c *client) GetResources(build Build) (Run, error) {

	response, err := c.get(build.URL + "/resources")
	if err != nil {
		return Run{}, err
	}

	inputs := Resource{}
	err = json.Unmarshal(response, &inputs)
	if err != nil {
		return Run{}, err
	}
	run := Run{
		Status:    build.Status,
		Resources: inputs,
	}

	return run, nil
}
