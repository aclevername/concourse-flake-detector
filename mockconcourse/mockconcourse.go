package mockconcourse

import "github.com/pivotal-cf/on-demand-service-broker/mockhttp"

func New() *mockhttp.Server {
	return mockhttp.StartServer("mock-concourse")
}
