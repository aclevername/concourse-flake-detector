package historybuilder_test

import (
	"github.com/aclevername/concourse-flake-detector/api"
	"github.com/aclevername/concourse-flake-detector/api/clientfake"
	"github.com/aclevername/concourse-flake-detector/historybuilder"

	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const gitResourcewithVersion = `
			{
	"inputs": [
		{
			"name": "concourse",
			"resource": "concourse",
			"type": "git",
			"version": {
				"ref": "%s"
			},
		  "pipeline_id": 1
		}
	]
}
`

var _ = Describe("historybuilder", func() {
	Describe("GetJobHistory", func() {
		It("gets a history of the job", func() {
			client := new(clientfake.FakeClient)
			testJob := api.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}

			buildList := `[{"status":"succeeded","api_url":"/api/v1/builds/1"},{"status":"failed","api_url":"/api/v1/builds/2"}]`
			client.GetReturnsOnCall(0, []byte(buildList), nil)

			client.GetReturnsOnCall(1, []byte(fmt.Sprintf(gitResourcewithVersion, "version1")), nil)

			client.GetReturnsOnCall(2, []byte(fmt.Sprintf(gitResourcewithVersion, "version1")), nil)

			history, err := historybuilder.GetJobHistory(client, testJob)

			Expect(err).NotTo(HaveOccurred())
			Expect(client.GetCallCount()).To(Equal(3))
			Expect(client.GetArgsForCall(0)).To(Equal("api/v1/teams/main/pipelines/main/jobs/fly/builds"))
			Expect(client.GetArgsForCall(1)).To(Equal("/api/v1/builds/1/resources"))
			Expect(client.GetArgsForCall(2)).To(Equal("/api/v1/builds/2/resources"))
			Expect(len(history)).To(Equal(2))
			Expect(history[0].Status).To(Equal("succeeded"))
			Expect(history[0].Resources.Inputs[0]).To(Equal(historybuilder.Input{
				Name:     "concourse",
				Resource: "concourse",
				Type:     "git",
				Version: historybuilder.Ref{
					Ref: "version1",
				},
				PipelineID: 1,
			}))

			Expect(history[1].Status).To(Equal("failed"))
			Expect(history[1].Resources.Inputs[0]).To(Equal(historybuilder.Input{
				Name:     "concourse",
				Resource: "concourse",
				Type:     "git",
				Version: historybuilder.Ref{
					Ref: "version1",
				},
				PipelineID: 1,
			}))
		})

		Context("when getting the job builds fails", func() {
			It("returns an error", func() {
				client := new(clientfake.FakeClient)
				testJob := api.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}

				buildList := `[{"status":"succeeded","api_url":"/api/v1/builds/1"},{"status":"failed","api_url":"/api/v1/builds/2"}]`
				client.GetReturnsOnCall(0, []byte(buildList), nil)
				client.GetReturnsOnCall(1, []byte{}, errors.New("failed"))

				_, err := historybuilder.GetJobHistory(client, testJob)
				Expect(err).To(MatchError("failed"))
			})
		})

		Context("when getting the a builds resources fails", func() {
			It("returns an error", func() {
				client := new(clientfake.FakeClient)
				testJob := api.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}

				client.GetReturnsOnCall(0, []byte{}, errors.New("failed"))

				_, err := historybuilder.GetJobHistory(client, testJob)
				Expect(err).To(MatchError("failed"))
			})
		})
	})
})
