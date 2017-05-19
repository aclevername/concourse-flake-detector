package main

import (
	"flag"
	"fmt"

	"net/http"

	"bytes"

	"os"

	"strconv"

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

	client := concourse.NewClient(func(url string) ([]byte, error) {
		fmt.Println("----------------------------\nSTART\n----------------------------")
		fmt.Printf("URL: %s\n", url)
		//response, err := http.Get(url)
		//if err != nil {
		//	return nil, err
		//}
		//buffer := new(bytes.Buffer)
		//buffer.ReadFrom(response.Body)
		//
		////fmt.Printf("RESPONSE: %s\n", string(buffer.Bytes()))
		////fmt.Println("----------------------------\nEND\n----------------------------")
		//return buffer.Bytes(), err

		var bearer = "Bearer " + *bearer
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("authorization", bearer)

		client := &http.Client{}

		response, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		buffer := new(bytes.Buffer)
		buffer.ReadFrom(response.Body)

		//fmt.Printf("RESPONSE: %s\n", string(buffer.Bytes()))
		//fmt.Println("----------------------------\nEND\n----------------------------")
		return buffer.Bytes(), err
	}, *url, *team)

	pipeline, err := client.GetPipeline(*name)

	//results := map[string]flake{}
	results := make([][]string, 0)

	if err != nil {
		panic(err)
	}
	for _, job := range pipeline.Jobs() {

		jobHistory, err := historybuilder.GetJobHistory(client, job, *count)
		if err != nil {
			panic(err)
		}

		jobFlakeCount, err := flakedetector.Detect(jobHistory)
		if err != nil {
			panic(err)
		}

		results = append(results, []string{job.Name, strconv.Itoa(len(jobHistory)), strconv.Itoa(jobFlakeCount)})
		//results[job.Name] = flake{
		//	total: len(jobHistory),
		//	count: jobFlakeCount,
		//}
		//fmt.Printf("\n----Result-----\nPipeline: %s\n", *name)
		//fmt.Printf("Job: %s, total runs: %d, flakeyness: %d\n", job.Name, len(jobHistory), jobFlakeCount)
	}

	//for jobName, result := range results {
	//	fmt.Printf("Job: %s, total runs: %d, flakeyness: %d\n", jobName, result.total, result.count)
	//}
	fmt.Println(results)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "builds", "flake"})

	for _, v := range results {
		table.Append(v)
	}
	table.Render() // Send output
}

type flake struct {
	count int
	total int
}
