package main

import (
	"flag"
	"fmt"
	"github.com/aclevername/concourse-flake-detector/api"
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
	var client api.Client
	client = new(httpclient.Client)

	pipeline, _ := api.GetPipeline(*url, *name, client)
	fmt.Printf("\n----Result-----\nPipeline: %s\n", *name)
	fmt.Printf("Job: %s, flakeyness: \n", pipeline.Jobs()[0].Name)
}
