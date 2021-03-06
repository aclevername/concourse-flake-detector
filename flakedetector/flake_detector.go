package flakedetector

import (
	"github.com/aclevername/concourse-flake-detector/concourse"
)

type result struct {
	passed      bool
	failedCount int
}

func Detect(runs []concourse.Run) (int, error) {
	resourceMap := map[string]*result{}
	for _, value := range runs {
		inputKey := inputArrayToString(value.Resources.Inputs)
		if _, ok := resourceMap[inputKey]; !ok {
			if value.Status == "failed" {
				resourceMap[inputKey] = &result{failedCount: 1, passed: false}
			} else {
				resourceMap[inputKey] = &result{failedCount: 0, passed: true}
			}
		} else {
			if value.Status == "failed" {
				resourceMap[inputKey].failedCount++
			} else {
				resourceMap[inputKey].passed = true
			}
		}
	}

	flakeCount := 0
	for _, value := range resourceMap {
		if value.passed {
			flakeCount += value.failedCount
		}
	}
	return flakeCount, nil
}

func inputArrayToString(input []concourse.Input) string {
	var content string
	for _, value := range input {
		content += value.Version.Ref
	}
	return content
}
