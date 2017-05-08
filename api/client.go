package api

//go:generate counterfeiter -o clientfake/fake_client.go . Client
type Client interface {
	Get(string) ([]byte, error)
}
