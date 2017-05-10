package flakedetector_test

import (
	"github.com/aclevername/concourse-flake-detector/flakedetector"
	"github.com/aclevername/concourse-flake-detector/historybuilder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("flakedetector", func() {
	Context("when given a list of runs", func() {
		It("returns the number of times a run has passed/failed with the same resource input", func() {

			runOne := historybuilder.Run{
				Status: "failed",
				Resources: historybuilder.Resource{
					Inputs: []historybuilder.Input{
						historybuilder.Input{
							Version: historybuilder.Ref{
								Ref: "version1",
							},
						},
					},
				},
			}
			runTwo := historybuilder.Run{
				Status: "succeeded",
				Resources: historybuilder.Resource{
					Inputs: []historybuilder.Input{
						historybuilder.Input{
							Version: historybuilder.Ref{
								Ref: "version1",
							},
						},
					},
				},
			}
			runThree := historybuilder.Run{
				Status: "failed",
				Resources: historybuilder.Resource{
					Inputs: []historybuilder.Input{
						historybuilder.Input{
							Version: historybuilder.Ref{
								Ref: "version1",
							},
						},
					},
				},
			}
			runs := []historybuilder.Run{
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

// given a list of runs, add each run to a map where---  key:runInputs, value: number of passes, number of failures
// Condition: if passes != 0 && failures != 0 then we have a flakey test.

// we pass into the flakedetector the Runs and something that does what the condition is
