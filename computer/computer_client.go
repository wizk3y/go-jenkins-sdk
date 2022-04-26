package computer

import (
	"encoding/json"
	"fmt"
	"net/http"

	restclient "github.com/wizk3y/go-jenkins-sdk/rest"
)

// ComputerInterface --
type ComputerInterface interface {
	CreateComputer(req *ComputerRequest) (bool, error)
	GetComputers() (*ComputersResponse, error)
	GetComputer(name string) (*Computer, error)
	UpdateComputer(req *ComputerRequest) (bool, error)
	DeleteComputer(name string) (bool, error)
}

type computerClient struct {
	client *restclient.RESTClient
}

// CreateComputer --
func (c *computerClient) CreateComputer(req *ComputerRequest) (bool, error) {
	// prepare option
	opts := make([]restclient.HttpOptionFunc, 0)
	d, _ := json.Marshal(req)
	opts = append(opts, restclient.HttpQueryOptionFunc("name", req.Name))
	opts = append(opts, restclient.HttpQueryOptionFunc("type", "hudson.slaves.DumbSlave$DescriptorImpl"))
	opts = append(opts, restclient.HttpQueryOptionFunc("json", string(d)))
	opts = append(opts, restclient.HttpHeaderOptionFunc("Content-Type", "application/x-www-form-urlencoded"))

	// do request
	resp, err := c.client.Do(http.MethodPost, "computer/doCreateItem", opts...)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// GetComputers --
func (c *computerClient) GetComputers() (*ComputersResponse, error) {
	// prepare option
	opts := make([]restclient.HttpOptionFunc, 0)
	opts = append(opts, restclient.HttpHeaderOptionFunc("Content-Type", "application/json"))

	// do request
	resp, err := c.client.Do(http.MethodGet, "computer/api/json", opts...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get output
	out := ComputersResponse{}
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// GetComputer --
func (c *computerClient) GetComputer(name string) (*Computer, error) {
	// default get master
	if len(name) == 0 {
		name = "(master)"
	}

	// prepare option
	opts := make([]restclient.HttpOptionFunc, 0)
	opts = append(opts, restclient.HttpHeaderOptionFunc("Content-Type", "application/json"))

	// do request
	resp, err := c.client.Do(http.MethodGet, fmt.Sprintf("computer/%v/api/json", name), opts...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get output
	out := Computer{}
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *computerClient) UpdateComputer(req *ComputerRequest) (bool, error) {
	// prepare option
	opts := make([]restclient.HttpOptionFunc, 0)
	d, _ := json.Marshal(req)
	opts = append(opts, restclient.HttpQueryOptionFunc("name", req.Name))
	opts = append(opts, restclient.HttpQueryOptionFunc("json", string(d)))
	opts = append(opts, restclient.HttpHeaderOptionFunc("Content-Type", "application/x-www-form-urlencoded"))

	// do request
	resp, err := c.client.Do(http.MethodPost, fmt.Sprintf("computer/%v/configSubmit", req.Name), opts...)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return true, nil
}

// DeleteComputer --
func (c *computerClient) DeleteComputer(name string) (bool, error) {
	// must have name
	if len(name) == 0 {
		return false, fmt.Errorf("missing computer name")
	}

	// prepare option
	opts := make([]restclient.HttpOptionFunc, 0)
	opts = append(opts, restclient.HttpHeaderOptionFunc("Content-Type", "application/x-www-form-urlencoded"))

	// do request
	resp, err := c.client.Do(http.MethodPost, fmt.Sprintf("computer/%v/doDelete", name), opts...)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// NewForConfig --
func NewForConfig(c *restclient.Config) (*computerClient, error) {
	client, err := c.Build()
	if err != nil {
		return nil, err
	}

	return &computerClient{client: client}, nil
}

// NewForConfigOrDie --
func NewForConfigOrDie(c *restclient.Config) *computerClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}

	return client
}

// New --
func New(client *restclient.RESTClient) *computerClient {
	return &computerClient{
		client: client,
	}
}
