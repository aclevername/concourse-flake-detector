package flakedetector

import "github.com/aclevername/concourse-flake-detector/historybuilder"

type result struct {
	passedCount int
	failedCount int
}

func Detect(runs []historybuilder.Run) (int, error) {
	resourceMap := map[[]historybuilder.Input]result{}
	for _, value := range runs {
		// if my map doesn't contain the runs resource input then add it and that pass/fail state
		if _, ok := resourceMap[value.Resources.Inputs]; !ok {
			//do something here
			if value.Status == "failed" {
				resourceMap[value.Resources.Inputs] = result{passedCount: 1}
			} else {
				resourceMap[value.Resources.Inputs] = result{failedCount: 1}
			}
		} else {
			if value.Status == "failed" {
				resourceMap[value.Resources.Inputs].passedCount++
			} else {
				resourceMap[value.Resources.Inputs].failedCount++
			}
		}
	}
	flakeCount := 0
	for _, value := range resourceMap {
		if value.failedCount != 0 && value.passedCount != 0 {
			flakeCount++
		}
	}

	return flakeCount, nil
}
