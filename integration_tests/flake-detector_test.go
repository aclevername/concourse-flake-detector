package integration_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os/exec"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"io"
	"time"
	"github.com/aclevername/concourse-flake-detector/mockconcourse"
	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
)

var _ = Describe("flake-detector", func(){
	var concourse *mockhttp.Server

	BeforeEach(func(){
		concourse = mockconcourse.New()
	})

	AfterEach(func(){
		concourse.VerifyMocks()
		concourse.Close()
	})
	It("lists the jobs", func(){
		jobName := "test-job"
		pipelineName := "test-pipeline"
		concourse.AppendMocks(mockconcourse.JobsForPipeline(pipelineName).RespondsWithJob(jobName, "test-job-url"))
		params := []string{"-url", concourse.URL, "-pipeline", pipelineName}
		_, logBuffer := runFlakeDetector(params...)

		Expect(logBuffer).To(gbytes.Say("Pipeline: %s", pipelineName))
		Expect(logBuffer).To(gbytes.Say("Job: %s, flakeyness: ", jobName))

	})
})

func runFlakeDetector(params ...string) (*gexec.Session, *gbytes.Buffer) {
	cmd := exec.Command(deleter, params...)
	logBuffer := gbytes.NewBuffer()
	session, err := gexec.Start(cmd, io.MultiWriter(GinkgoWriter, logBuffer), GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session, 10*time.Second).Should(gexec.Exit(0))
	return session, logBuffer
}