package restclient

import (
	"fmt"
	"net/http"
)

// RESTClient --
type RESTClient struct {
	baseURL string

	client *http.Client
}

func (c *RESTClient) Do(method, path string, opts ...HttpOptionFunc) (*http.Response, error) {
	// create request
	fullUrl := fmt.Sprintf("%v/%v", c.baseURL, path)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return nil, err
	}

	// apply option
	for _, opt := range opts {
		opt.Apply(req)
	}

	// do request
	var resp *http.Response
	resp, err = c.client.Do(req)

	return resp, err
}
