package pipeline_test

import (
	"errors"
	"github.com/aclevername/concourse-flake-detector/pipeline"
	fakes "github.com/aclevername/concourse-flake-detector/pipeline/pipelinefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pipeline", func() {
	Describe("New", func() {
		var client *fakes.FakeHTTPClient

		It("returns the pipeline", func() {
			client = new(fakes.FakeHTTPClient)
			response := `[{"name":"job0","url":"/foo/bar"}, {"name":"job1","url":"/bar/baz"}, {"name":"job2","url":"/baz/foo"}]`
			client.GetReturns([]byte(response), nil)

			testPipeline, err := pipeline.New("test-concourse.com", "test-pipeline", client)

			Expect(err).NotTo(HaveOccurred())

			By("calling the concourse endpoint")
			Expect(client.GetCallCount()).To(Equal(1))
			Expect(client.GetArgsForCall(0)).To(Equal("test-concourse.com/api/v1/pipelines/test-pipeline/jobs"))

			jobs := testPipeline.Jobs()

			By("containing the list of jobs")
			Expect(len(jobs)).To(Equal(3))
			Expect(jobs[0].Name).To(Equal("job0"))
			Expect(jobs[0].URL).To(Equal("/foo/bar"))
			Expect(jobs[1].Name).To(Equal("job1"))
			Expect(jobs[1].URL).To(Equal("/bar/baz"))
			Expect(jobs[2].Name).To(Equal("job2"))
			Expect(jobs[2].URL).To(Equal("/baz/foo"))

		})

		Context("when the get request fails", func() {
			It("returns an error", func() {
				client = new(fakes.FakeHTTPClient)
				client.GetReturns(nil, errors.New("I failed"))

				_, err := pipeline.New("test-concourse.com", "test-pipeline", client)

				Expect(err).To(MatchError("I failed"))
			})
		})

		Context("when the reponse returns invalid json", func() {
			It("returns the error", func() {
				client = new(fakes.FakeHTTPClient)
				response := `defo not json`
				client.GetReturns([]byte(response), nil)

				_, err := pipeline.New("test-concourse.com", "test-pipeline", client)

				Expect(err).To(HaveOccurred())
			})
		})

	})
})
