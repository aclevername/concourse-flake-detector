package concourse_test

import (
	"errors"

	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/concourse/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pipeline", func() {
	Describe("GetPipeline", func() {
		var (
			fakeGet      *fake.FakeGetter
			client       concourse.ClientInterface
			testPipeline concourse.Pipeline
			err          error
		)

		BeforeEach(func() {
			fakeGet = new(fake.FakeGetter)
			client = concourse.NewClient(fakeGet.Spy, "example.com", "")
		})

		JustBeforeEach(func() {
			testPipeline, err = client.GetPipeline("test-pipeline")
		})

		Context("when the get request succeeds", func() {
			BeforeEach(func() {
				response := `[{"name":"job0","url":"/foo/bar"}, {"name":"job1","url":"/bar/baz"}, {"name":"job2","url":"/baz/foo"}]`
				fakeGet.Returns([]byte(response), nil)
			})

			It("returns the pipeline", func() {

				Expect(err).NotTo(HaveOccurred())

				By("calling the concourse endpoint")
				Expect(fakeGet.CallCount()).To(Equal(1))
				Expect(fakeGet.ArgsForCall(0)).To(Equal("example.com/api/v1/pipelines/test-pipeline/jobs"))

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
				fakeGet.Returns(nil, errors.New("I failed"))
			})

			It("returns an error", func() {
				Expect(err).To(MatchError("I failed"))
			})
		})

		Context("when the reponse returns invalid json", func() {
			BeforeEach(func() {
				response := `defo not json`
				fakeGet.Returns([]byte(response), nil)
			})

			It("returns the error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
