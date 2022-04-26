package jenkins

import (
	"github.com/wizk3y/go-jenkins-sdk/computer"
	restclient "github.com/wizk3y/go-jenkins-sdk/rest"
)

// Interface --
type Interface interface {
	Computer() computer.ComputerInterface
}

// Clientset --
type Clientset struct {
	computerClient computer.ComputerInterface
}

// Computer --
func (c *Clientset) Computer() computer.ComputerInterface {
	return c.computerClient
}

// NewForConfig --
func NewForConfig(c *restclient.Config) (*Clientset, error) {
	var (
		cs  Clientset
		err error
	)

	cs.computerClient, err = computer.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}

// NewForConfigOrDie --
func NewForConfigOrDie(c *restclient.Config) *Clientset {
	var cs Clientset

	cs.computerClient = computer.NewForConfigOrDie(c)

	return &cs
}

// New --
func New(client *restclient.RESTClient) *Clientset {
	var cs Clientset

	cs.computerClient = computer.New(client)

	return &cs
}
