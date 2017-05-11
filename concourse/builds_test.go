package concourse_test

import (
	"errors"

	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/concourse/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetBuilds", func() {

	It("returns the builds for a job", func() {
		fakeGet := new(fake.FakeGetter)
		client := concourse.NewClient(fakeGet.Spy, "example.com", "")

		buildList := `[{"status":"succeeded","api_url":"/concourse/v1/builds/1"},{"status":"failed","api_url":"/concourse/v1/builds/2"}]`
		fakeGet.Returns([]byte(buildList), nil)
		builds, err := client.GetBuilds(concourse.Job{URL: "example.com/api/v1/foo/job"})
		Expect(err).NotTo(HaveOccurred())
		Expect(fakeGet.ArgsForCall(0)).To(Equal("example.com/api/v1/foo/job/builds"))
		Expect(builds).To(Equal([]concourse.Build{
			{
				URL:    "/concourse/v1/builds/1",
				Status: "succeeded",
			},
			{
				URL:    "/concourse/v1/builds/2",
				Status: "failed",
			},
		}))
	})

	Context("when the get request fails", func() {
		It("returns an error", func() {
			fakeGet := new(fake.FakeGetter)
			client := concourse.NewClient(fakeGet.Spy, "example.com", "")
			fakeGet.Returns([]byte(""), errors.New("failed"))
			_, err := client.GetBuilds(concourse.Job{URL: "example.com"})
			Expect(err).To(MatchError("failed"))
		})
	})

	Context("when the get returns invalid json", func() {
		It("returns an error", func() {
			fakeGet := new(fake.FakeGetter)
			client := concourse.NewClient(fakeGet.Spy, "example.com", "")
			fakeGet.Returns([]byte("not valid json"), nil)
			_, err := client.GetBuilds(concourse.Job{URL: "example.com"})
			Expect(err).To(HaveOccurred())
		})
	})
})
