package api_test

import (
	"errors"
	"github.com/aclevername/concourse-flake-detector/api"
	"github.com/aclevername/concourse-flake-detector/api/clientfake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pipeline", func() {
	Describe("GetPipeline", func() {
		var (
			client       *clientfake.FakeClient
			testPipeline api.Pipeline
			err          error
		)

		BeforeEach(func() {
			client = new(clientfake.FakeClient)
		})

		JustBeforeEach(func() {
			testPipeline, err = api.GetPipeline("test-concourse.com", "test-pipeline", client)
		})

		Context("when the get request succeeds", func() {
			BeforeEach(func() {
				response := `[{"name":"job0","api_url":"/foo/bar"}, {"name":"job1","api_url":"/bar/baz"}, {"name":"job2","api_url":"/baz/foo"}]`
				client.GetReturns([]byte(response), nil)
			})
			It("returns the pipeline", func() {

				Expect(err).NotTo(HaveOccurred())

				By("calling the concourse endpoint")
				Expect(client.GetCallCount()).To(Equal(1))
				Expect(client.GetArgsForCall(0)).To(Equal("api/v1/pipelines/test-pipeline/jobs"))

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
		})

		Context("when the get request fails", func() {
			BeforeEach(func() {
				client.GetReturns(nil, errors.New("I failed"))
			})

			It("returns an error", func() {
				Expect(err).To(MatchError("I failed"))
			})
		})

		Context("when the reponse returns invalid json", func() {
			BeforeEach(func() {
				response := `defo not json`
				client.GetReturns([]byte(response), nil)
			})

			It("returns the error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
