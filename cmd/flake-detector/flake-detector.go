package main

import (
	"bytes"
	"flag"
	"fmt"

	"net/http"

	"github.com/aclevername/concourse-flake-detector/concourse"
	"github.com/aclevername/concourse-flake-detector/flakedetector"
	"github.com/aclevername/concourse-flake-detector/historybuilder"
)

func main() {
	url := flag.String("url", "", "concourse url")
	name := flag.String("pipeline", "", "pipeline name")
	team := flag.String("team", "", "team name, optional")
	count := flag.Int("count", 0, "how many of the latest builds to scan through, optional")

	flag.Parse()
	fmt.Printf("configuration: url %s, pipeline %s\n", *url, *name)
	if *url == "" || *name == "" {
		panic("please configure correctly using -url and -pipeline")
	}

	client := concourse.NewClient(func(url string) ([]byte, error) {
		fmt.Println("----------------------------\nSTART\n----------------------------")
		fmt.Printf("URL: %s\n", url)
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(response.Body)

		//fmt.Printf("RESPONSE: %s\n", string(buffer.Bytes()))
		//fmt.Println("----------------------------\nEND\n----------------------------")
		return buffer.Bytes(), err
	}, *url, *team)

	pipeline, err := client.GetPipeline(*name)

	if err != nil {
		panic(err)
	}

	jobHistory, err := historybuilder.GetJobHistory(client, pipeline.Jobs()[0], *count)
	if err != nil {
		panic(err)
	}

	jobFlakeCount, err := flakedetector.Detect(jobHistory)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n----Result-----\nPipeline: %s\n", *name)
	fmt.Printf("Job: %s, flakeyness: %d\n", pipeline.Jobs()[0].Name, jobFlakeCount)

}
