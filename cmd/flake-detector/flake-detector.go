package main

import (
	"flag"
	"fmt"
	"github.com/aclevername/concourse-flake-detector/api"
	"github.com/aclevername/concourse-flake-detector/flakedetector"
	"github.com/aclevername/concourse-flake-detector/historybuilder"
	"github.com/aclevername/concourse-flake-detector/httpclient"
)

func main() {
	url := flag.String("url", "", "concourse url")
	name := flag.String("pipeline", "", "pipeline name")
	flag.Parse()
	fmt.Printf("configuration: url %s, pipeline %s\n", *url, *name)
	if *url == "" || *name == "" {
		panic("please configure correctly using -url and -pipeline")
	}

	client := &realClient{
		BaseURL: *url,
	}

	pipeline, _ := api.GetPipeline(*url, *name, client)

	fmt.Println("------------------")

	fmt.Println(pipeline.Jobs()[0].URL)
	fmt.Println(pipeline.Jobs()[0].Name)
	jobHistory, err := historybuilder.GetJobHistory(client, pipeline.Jobs()[0])
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

type realClient struct {
	BaseURL string
}

func (rc *realClient) Get(url string) ([]byte, error) {
	var client api.Client
	client = new(httpclient.Client)
	return client.Get(rc.BaseURL + "/" + url)
}
