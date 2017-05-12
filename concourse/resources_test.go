package concourse_test

import (
	"errors"
	"fmt"

	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/concourse/fake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetResources", func() {

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

	It("returns the resources for a build", func() {
		fakeGet := new(fake.FakeGetter)
		client := concourse.NewClient(fakeGet.Spy, "example.com", "")

		fakeGet.ReturnsOnCall(0, []byte(fmt.Sprintf(gitResourcewithVersion, "version1")), nil)
		fakeGet.ReturnsOnCall(1, []byte(fmt.Sprintf(gitResourcewithVersion, "version1")), nil)

		run, err := client.GetResources(concourse.Build{
			Status: "succeeded",
			URL:    "/builds",
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(fakeGet.ArgsForCall(0)).To(Equal("example.com/builds/resources"))
		Expect(run).To(Equal(concourse.Run{
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
		}))
	})

	Context("when the get request fails", func() {
		It("returns an error", func() {
			fakeGet := new(fake.FakeGetter)
			client := concourse.NewClient(fakeGet.Spy, "fake.com", "")
			fakeGet.Returns([]byte(""), errors.New("failed"))
			_, err := client.GetResources(concourse.Build{URL: "example.com/builds"})
			Expect(err).To(MatchError("failed"))
		})
	})

	Context("when the get returns invalid json", func() {
		It("returns an error", func() {
			fakeGet := new(fake.FakeGetter)
			client := concourse.NewClient(fakeGet.Spy, "fake.com", "")
			fakeGet.Returns([]byte("not valid json"), nil)
			_, err := client.GetResources(concourse.Build{URL: "example.com/builds"})
			Expect(err).To(HaveOccurred())
		})
	})
})
