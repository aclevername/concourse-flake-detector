package main

import (
	"flag"
	"fmt"

	"net/http"

	"bytes"

	"os"

	"strconv"

	"io"

	"crypto/tls"

	"strings"

	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/flakedetector"
	"github.com/aclevername/concourse-flake-detector/historybuilder"
	"github.com/olekukonko/tablewriter"
)

func main() {
	url, pipelineName, team, count, bearer, debug, skipTls := initialiseFlags()

	flag.Parse()

	checkConfiguredCorrectly(url, pipelineName)

	getFunc := intitaliseBearer(team, url, bearer, debug, skipTls)

	client := concourse.NewClient(getFunc, *url, *team)

	pipeline, err := client.GetPipeline(*pipelineName)

	if err != nil {
		exitWithError(err)
	}

	results := scanConcourse(pipeline, client, count)

	print(results, *pipelineName)
}

func checkConfiguredCorrectly(url *string, pipelineName *string) {
	fmt.Printf("\n\nconfiguration: url %s, pipeline %s\n", *url, *pipelineName)
	if *url == "" || *pipelineName == "" {
		exitWithError(fmt.Errorf("please configure correctly using -url and -pipeline"))
	}
}

func scanConcourse(pipeline concourse.Pipeline, client concourse.ClientInterface, count *int) [][]string {
	results := make([][]string, 0)
	for _, job := range pipeline.Jobs() {

		jobHistory, err := historybuilder.GetJobHistory(client, job, *count)
		if err != nil {
			exitWithError(err)
		}

		jobFlakeCount, err := flakedetector.Detect(jobHistory)
		if err != nil {
			exitWithError(err)
		}

		results = append(results, []string{job.Name, strconv.Itoa(len(jobHistory)), strconv.Itoa(jobFlakeCount)})
	}
	return results
}

func intitaliseBearer(team *string, url *string, bearer *string, debug *bool, skipTls *bool) func(url string) ([]byte, error) {
	var bearerURL string
	if *team != "" {
		bearerURL = fmt.Sprintf("%s/api/v1/teams/%s/auth/token", *url, *team)
	} else {
		bearerURL = fmt.Sprintf("%s/api/v1/auth/token", *url)
	}
	var getFunc func(url string) ([]byte, error)
	if *bearer != "" {
		getFunc = clientWithBearer(*bearer, bearerURL, *debug, *skipTls)
	} else {
		getFunc = client(bearerURL, *debug)

	}
	return getFunc
}

func initialiseFlags() (*string, *string, *string, *int, *string, *bool, *bool) {
	url := flag.String("url", "", "concourse url")
	pipelineName := flag.String("pipeline", "", "pipeline pipelineName")
	team := flag.String("team", "", "team name, optional")
	count := flag.Int("count", 0, "how many of the latest builds to scan through, optional")
	bearer := flag.String("bearer", "", "bearer token")
	debug := flag.Bool("debug", false, "debug flag")
	skipTls := flag.Bool("insecure-tls", false, "TLS accepts any certificate presented by the server and any host name in that certificate")
	return url, pipelineName, team, count, bearer, debug, skipTls
}

type flake struct {
	count int
	total int
}

func getBody(body io.ReadCloser) []byte {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	return buffer.Bytes()
}

func checkAuth(body []byte, bearerAddress string) error {
	if bytes.Contains(body, []byte("not authorized")) {
		return fmt.Errorf("Please provide a bearer token using the -bearer flag, obtain the token by logging into: %s", bearerAddress)
	}
	return nil
}

func checkInvalidTLS(err error) error {
	if strings.Contains(err.Error(), "certificate signed by unknown authority") {
		return fmt.Errorf("It appears your pipeline hasn't configured TLS correctly, in order to proceed add the -insecure-tls flag.")
	}
	return err
}

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func clientWithBearer(bearer, bearerURL string, debug, skipTLS bool) func(url string) ([]byte, error) {
	return func(url string) ([]byte, error) {
		if debug {
			fmt.Println("Get: " + url)
		}
		var bearer = "Bearer " + bearer
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("authorization", bearer)

		var client *http.Client
		if skipTLS {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client = &http.Client{Transport: tr}
		} else {
			client = &http.Client{}

		}
		response, err := client.Do(req)
		if err != nil {
			return nil, checkInvalidTLS(err)
		}
		defer response.Body.Close()

		body := getBody(response.Body)

		return body, checkAuth(body, bearerURL)

	}
}

func client(bearerURL string, debug bool) func(url string) ([]byte, error) {
	return func(url string) ([]byte, error) {
		if debug {
			fmt.Println("Get: " + url)
		}
		response, err := http.Get(url)
		if err != nil {
			return nil, checkInvalidTLS(err)
		}
		body := getBody(response.Body)

		return body, checkAuth(body, bearerURL)
	}
}

func print(results [][]string, pipeline string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "builds", "flakes"})

	for _, v := range results {
		table.Append(v)
	}
	fmt.Println("Pipeline: " + pipeline)
	table.Render()
}
