package historybuilder_test

import (
	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/historybuilder"

	"errors"

	"github.com/aclevername/concourse-flake-detector/concourse/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("historybuilder", func() {
	Describe("GetJobHistory", func() {
		It("gets a history of the job", func() {
			client := new(fake.FakeClientInterface)
			testJob := concourse.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}

			client.GetBuildsReturns([]concourse.Build{
				{
					URL:    "/concourse/v1/builds/1",
					Status: "succeeded",
				},
				{
					URL:    "/concourse/v1/builds/2",
					Status: "failed",
				},
			}, nil)

			client.GetResourcesReturnsOnCall(0, concourse.Run{
				Status: "succeeded",
				Resources: concourse.Resource{
					Inputs: []concourse.Input{
						{
							Name:     "concourse",
							Resource: "concourse",
							Type:     "git",
							Version: concourse.Ref{
								Ref: "version1",
							},
							PipelineID: 1,
						},
					},
				},
			}, nil)

			client.GetResourcesReturnsOnCall(1, concourse.Run{
				Status: "failed",
				Resources: concourse.Resource{
					Inputs: []concourse.Input{
						{
							Name:     "concourse",
							Resource: "concourse",
							Type:     "git",
							Version: concourse.Ref{
								Ref: "version1",
							},
							PipelineID: 1,
						},
					},
				},
			}, nil)

			history, err := historybuilder.GetJobHistory(client, testJob, 0)

			Expect(err).NotTo(HaveOccurred())
			Expect(history[0].Status).To(Equal("succeeded"))
			Expect(history[0].Resources.Inputs[0]).To(Equal(concourse.Input{
				Name:     "concourse",
				Resource: "concourse",
				Type:     "git",
				Version: concourse.Ref{
					Ref: "version1",
				},
				PipelineID: 1,
			}))

			Expect(history[1].Status).To(Equal("failed"))
			Expect(history[1].Resources.Inputs[0]).To(Equal(concourse.Input{
				Name:     "concourse",
				Resource: "concourse",
				Type:     "git",
				Version: concourse.Ref{
					Ref: "version1",
				},
				PipelineID: 1,
			}))
		})

		Context("when getting the job builds fails", func() {
			It("returns an error", func() {
				//client := new(clientfake.FakeClient)
				//testJob := concourse.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}
				//
				//buildList := `[{"status":"succeeded","api_url":"/concourse/v1/builds/1"},{"status":"failed","api_url":"/concourse/v1/builds/2"}]`
				//client.GetReturnsOnCall(0, []byte(buildList), nil)
				//client.GetReturnsOnCall(1, []byte{}, errors.New("failed"))

				client := new(fake.FakeClientInterface)
				testJob := concourse.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}

				client.GetBuildsReturns([]concourse.Build{}, errors.New("failed getting build"))

				_, err := historybuilder.GetJobHistory(client, testJob, 0)
				Expect(err).To(MatchError("failed getting build"))
			})
		})

		Context("when getting the a builds resources fails", func() {
			It("returns an error", func() {

				client := new(fake.FakeClientInterface)
				testJob := concourse.Job{Name: "test-job", URL: "/teams/main/pipelines/main/jobs/fly"}

				client.GetBuildsReturns([]concourse.Build{
					{
						URL:    "/concourse/v1/builds/1",
						Status: "succeeded",
					},
					{
						URL:    "/concourse/v1/builds/2",
						Status: "failed",
					},
				}, nil)

				client.GetResourcesReturns(concourse.Run{}, errors.New("failed getting resource"))

				_, err := historybuilder.GetJobHistory(client, testJob, 0)
				Expect(err).To(MatchError("failed getting resource"))
			})
		})
	})
})
