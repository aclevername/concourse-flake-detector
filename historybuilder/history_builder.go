package historybuilder

import (
	"github.com/aclevername/concourse-flake-detector/concourse"
)

func GetJobHistory(client concourse.ClientInterface, job concourse.Job, runCount int) ([]concourse.Run, error) {
	builds, err := client.GetBuilds(job)

	if runCount == 0 {
		runCount = len(builds)
	}

	if err != nil {
		return []concourse.Run{}, err
	}
	var history []concourse.Run

	count := 0
	for _, build := range builds {
		count++
		resources, err := client.GetResources(build)
		if err != nil {
			return []concourse.Run{}, err
		}
		history = append(history, resources)
		if count == runCount {
			return history, nil
		}
	}
	return history, nil
}
