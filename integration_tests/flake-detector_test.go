package integration_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/aclevername/concourse-flake-detector/mockconcourse"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/on-demand-service-broker/mockhttp"
)

var _ = Describe("flake-detector", func() {
	var concourse *mockhttp.Server

	BeforeEach(func() {
		concourse = mockconcourse.New()
	})

	AfterEach(func() {
		concourse.VerifyMocks()
		concourse.Close()
	})

	Context("when the job has 0 flakey run", func() {

		It("lists the jobs", func() {
			jobName := "test-job"
			pipelineName := "test-pipeline"

			buildList := fmt.Sprintf(`[{"status":"succeeded","api_url":"/api/v1/builds/1"},{"status":"failed","api_url":"/api/v1/builds/2"}]`)

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
			jobUrl := fmt.Sprintf("/pipelines/%s/jobs/%s", pipelineName, jobName)
			concourse.AppendMocks(
				mockconcourse.JobsForPipeline(pipelineName, "").RespondsWithJob(jobName, fmt.Sprintf("%s", jobUrl)),
				mockconcourse.BuildsForJob(jobUrl).RespondsWithBuilds(buildList),
				mockconcourse.ResourcesForBuild("/api/v1/builds/1").RespondsWith(fmt.Sprintf(gitResourcewithVersion, "v1")),
				mockconcourse.ResourcesForBuild("/api/v1/builds/2").RespondsWith(fmt.Sprintf(gitResourcewithVersion, "v2")),
			)
			params := []string{"-url", concourse.URL, "-pipeline", pipelineName}
			_, logBuffer := runFlakeDetector(0, params...)

			Expect(logBuffer).To(gbytes.Say("Pipeline: %s", pipelineName))
			Expect(string(logBuffer.Contents())).To(ContainSubstring("%s |      2 |     0 |", jobName))

		})
	})

	Context("when the job has 1 flakey run", func() {
		It("lists the jobs", func() {
			jobName := "test-job"
			pipelineName := "test-pipeline"

			buildList := fmt.Sprintf(`[{"status":"succeeded","api_url":"/api/v1/builds/1"},{"status":"failed","api_url":"/api/v1/builds/2"}]`)

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
			jobUrl := fmt.Sprintf("/pipelines/%s/jobs/%s", pipelineName, jobName)
			concourse.AppendMocks(
				mockconcourse.JobsForPipeline(pipelineName, "").RespondsWithJob(jobName, fmt.Sprintf("%s", jobUrl)),
				mockconcourse.BuildsForJob(jobUrl).RespondsWithBuilds(buildList),
				mockconcourse.ResourcesForBuild("/api/v1/builds/1").RespondsWith(fmt.Sprintf(gitResourcewithVersion, "v1")),
				mockconcourse.ResourcesForBuild("/api/v1/builds/2").RespondsWith(fmt.Sprintf(gitResourcewithVersion, "v1")),
			)
			params := []string{"-url", concourse.URL, "-pipeline", pipelineName}
			_, logBuffer := runFlakeDetector(0, params...)

			Expect(logBuffer).To(gbytes.Say("Pipeline: %s", pipelineName))
			Expect(string(logBuffer.Contents())).To(ContainSubstring("%s |      2 |     1 |", jobName))

		})

	})

	Context("when teams are configured", func() {
		It("lists the jobs", func() {
			jobName := "test-job"
			pipelineName := "test-pipeline"
			team := "foo"

			buildList := fmt.Sprintf(`[{"status":"succeeded","api_url":"/api/v1/builds/1"},{"status":"failed","api_url":"/api/v1/builds/2"}]`)

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
			//// example team URL api/v1/teams/main/pipelines/main/jobs/fly/builds

			jobUrl := fmt.Sprintf("/teams/foo/pipelines/%s/jobs/%s", pipelineName, jobName)
			concourse.AppendMocks(
				mockconcourse.JobsForPipeline(pipelineName, team).RespondsWithJob(jobName, fmt.Sprintf("%s", jobUrl)),
				mockconcourse.BuildsForJob(jobUrl).RespondsWithBuilds(buildList),
				mockconcourse.ResourcesForBuild("/api/v1/builds/1").RespondsWith(fmt.Sprintf(gitResourcewithVersion, "v1")),
				mockconcourse.ResourcesForBuild("/api/v1/builds/2").RespondsWith(fmt.Sprintf(gitResourcewithVersion, "v2")),
			)
			params := []string{"-url", concourse.URL, "-pipeline", pipelineName, "-team", team}
			_, logBuffer := runFlakeDetector(0, params...)

			Expect(logBuffer).To(gbytes.Say("Pipeline: %s", pipelineName))
			Expect(string(logBuffer.Contents())).To(ContainSubstring("%s |      2 |     0 |", jobName))

		})
	})

	Context("when the get request returns not authorized", func() {
		It("tells the user how to authenticate with the team", func() {
			pipelineName := "test-pipeline"
			team := "foo"

			concourse.AppendMocks(
				mockconcourse.JobsForPipeline(pipelineName, team).RespondsUnauthorizedWith("not authorized"),
			)
			params := []string{"-url", concourse.URL, "-pipeline", pipelineName, "-team", team}
			_, logBuffer := runFlakeDetector(1, params...)

			expectedMsg := fmt.Sprintf("Please provide a bearer token using the -bearer flag, obtain the token by logging into: %s/api/v1/teams/%s/auth/token", concourse.URL, team)
			Expect(logBuffer).To(gbytes.Say(expectedMsg))
		})

		It("tells the user how to authenticate with no team", func() {
			pipelineName := "test-pipeline"

			concourse.AppendMocks(
				mockconcourse.JobsForPipeline(pipelineName, "").RespondsUnauthorizedWith("not authorized"),
			)
			params := []string{"-url", concourse.URL, "-pipeline", pipelineName}
			_, logBuffer := runFlakeDetector(1, params...)

			expectedMsg := fmt.Sprintf("Please provide a bearer token using the -bearer flag, obtain the token by logging into: %s/api/v1/auth/token", concourse.URL)
			Expect(logBuffer).To(gbytes.Say(expectedMsg))
		})
	})

})

func runFlakeDetector(exitcode int, params ...string) (*gexec.Session, *gbytes.Buffer) {
	cmd := exec.Command(deleter, params...)
	logBuffer := gbytes.NewBuffer()
	session, err := gexec.Start(cmd, io.MultiWriter(GinkgoWriter, logBuffer), GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session, 10*time.Second).Should(gexec.Exit(exitcode))
	return session, logBuffer
}
