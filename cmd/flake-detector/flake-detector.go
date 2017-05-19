package main

import (
	"flag"
	"fmt"

	"net/http"

	"bytes"

	"os"

	"strconv"

	"io"

	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/flakedetector"
	"github.com/aclevername/concourse-flake-detector/historybuilder"
	"github.com/olekukonko/tablewriter"
)

func main() {
	url := flag.String("url", "", "concourse url")
	name := flag.String("pipeline", "", "pipeline name")
	team := flag.String("team", "", "team name, optional")
	count := flag.Int("count", 0, "how many of the latest builds to scan through, optional")
	bearer := flag.String("bearer", "", "bearer token")

	flag.Parse()
	fmt.Printf("configuration: url %s, pipeline %s\n", *url, *name)
	if *url == "" || *name == "" {
		panic("please configure correctly using -url and -pipeline")
	}
	var bearerURL string
	if *team != "" {
		bearerURL = fmt.Sprintf("%s/api/v1/teams/%s/auth/token", *url, *team)
	} else {
		bearerURL = fmt.Sprintf("%s/api/v1/auth/token", *url)

	}

	var getFunc func(url string) ([]byte, error)

	if *bearer != "" {
		getFunc = func(url string) ([]byte, error) {
			var bearer = "Bearer " + *bearer
			req, err := http.NewRequest("GET", url, nil)
			req.Header.Add("authorization", bearer)

			client := &http.Client{}

			response, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer response.Body.Close()

			body := getBody(response.Body)

			return body, checkAuth(body, bearerURL)

		}
	} else {
		getFunc = func(url string) ([]byte, error) {
			response, err := http.Get(url)
			if err != nil {
				return nil, err
			}
			body := getBody(response.Body)

			return body, checkAuth(body, bearerURL)
		}
	}

	client := concourse.NewClient(getFunc, *url, *team)

	pipeline, err := client.GetPipeline(*name)

	if err != nil {
		exitWithError(err)
	}

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

	fmt.Println(results)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "builds", "flake"})

	for _, v := range results {
		table.Append(v)
	}
	fmt.Println("Pipeline: " + *name)
	table.Render()
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

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
