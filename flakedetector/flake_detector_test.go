package flakedetector_test

import (
	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/flakedetector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("flakedetector", func() {
	Context("when given a list of runs", func() {
		It("returns the number of times a run has passed/failed with the same resource input", func() {

			runOne := concourse.Run{
				Status: "failed",
				Resources: concourse.Resource{
					Inputs: []concourse.Input{
						{
							Version: concourse.Ref{
								Ref: "version1",
							},
						},
					},
				},
			}
			runTwo := concourse.Run{
				Status: "succeeded",
				Resources: concourse.Resource{
					Inputs: []concourse.Input{
						{
							Version: concourse.Ref{
								Ref: "version1",
							},
						},
					},
				},
			}
			runThree := concourse.Run{
				Status: "failed",
				Resources: concourse.Resource{
					Inputs: []concourse.Input{
						{
							Version: concourse.Ref{
								Ref: "version1",
							},
						},
					},
				},
			}
			runs := []concourse.Run{
				runOne,
				runTwo,
				runThree,
			}
			occurences, err := flakedetector.Detect(runs)

			Expect(err).NotTo(HaveOccurred())
			Expect(occurences).To(Equal(2))
		})
	})
})
