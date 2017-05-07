package pipeline_test

import (
	"github.com/aclevername/concourse-flake-detector/pipeline"
	fakes "github.com/aclevername/concourse-flake-detector/pipeline/pipelinefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pipeline", func() {
	Describe("New", func() {
		var client       *fakes.FakeHTTPClient

		It("returns a list of jobs", func() {
			client = new(fakes.FakeHTTPClient)

			response := `[{"name":"job1"}, {"name":"job2"}, {"name":"job3"}]`

			client.GetReturns([]byte(response), nil)

			testPipeline := pipeline.New("test-concourse.com", "test-pipeline", client)
			jobs := testPipeline.Jobs()

			By("hitting the job url")
			Expect(client.GetCallCount()).To(Equal(1))

			Expect(len(jobs)).To(Equal(3))
		})

	})
})
