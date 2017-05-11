package historybuilder

import (
	"github.com/aclevername/concourse-flake-detector/concourse"
)

func GetJobHistory(client concourse.ClientInterface, job concourse.Job) ([]concourse.Run, error) {
	builds, err := client.GetBuilds(job)

	if err != nil {
		return []concourse.Run{}, err
	}
	var history []concourse.Run
	for _, build := range builds {
		resources, err := client.GetResources(build)
		if err != nil {
			return []concourse.Run{}, err
		}
		history = append(history, resources)
	}
	return history, nil
}
