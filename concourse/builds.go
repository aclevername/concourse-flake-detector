package concourse

import (
	"encoding/json"
	"fmt"
)

type Build struct {
	Status string `json:"status"`
	URL    string `json:"api_url"`
}

func (c *client) GetBuilds(job Job) ([]Build, error) {
	fmt.Println("+++++++++++++++++++++++++++++", job.URL)
	response, err := c.get(fmt.Sprintf("%s/builds", job.URL))
	if err != nil {
		return []Build{}, err
	}

	builds := make([]Build, 0)
	err = json.Unmarshal(response, &builds)
	if err != nil {
		return []Build{}, err
	}

	return builds, nil
}