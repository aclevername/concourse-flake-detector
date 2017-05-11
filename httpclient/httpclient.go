package httpclient

import (
	"bytes"
	"net/http"
)

type Client struct {
}

func (c *Client) Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)

	return buffer.Bytes(), err
}
