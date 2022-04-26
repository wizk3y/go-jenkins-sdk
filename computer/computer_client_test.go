package computer_test

import (
	"log"
	"testing"

	"github.com/wizk3y/go-jenkins-sdk/computer"
	restclient "github.com/wizk3y/go-jenkins-sdk/rest"
)

var (
	baseURL  = "http://localhost:8080"
	username = "example@example.com"
	password = ""
)

func Test_CreateComputer(t *testing.T) {
	config := restclient.NewConfig(baseURL, username, password)

	c, err := computer.NewForConfig(config)
	if err != nil {
		log.Println("error when create client", err)
		return
	}

	req := computer.ComputerRequest{
		Name:            "test_slave",
		NodeDescription: "",
		NumExecutors:    "4",
		RemoteFS:        "/var/lib/jenkins",
		LabelString:     "build_prod master",
		Mode:            "NORMAL", // "NORMAL"
		Launcher: computer.Launcher{
			StaplerClass:  "hudson.plugins.sshslaves.SSHLauncher",
			Class:         "hudson.plugins.sshslaves.SSHLauncher",
			Host:          "10.99.2.198",
			IncludeUser:   "false",
			CredentialsID: "slave_ssh",
			SSHHostKeyVerificationStrategy: computer.SSHHostKeyVerificationStrategy{
				StaplerClass: "hudson.plugins.sshslaves.verifiers.KnownHostsFileKeyVerificationStrategy",
				Class:        "hudson.plugins.sshslaves.verifiers.KnownHostsFileKeyVerificationStrategy",
			},
			Port:                 "16889",
			JavaPath:             "",
			JvmOptions:           "",
			PrefixStartSlaveCmd:  "",
			SuffixStartSlaveCmd:  "",
			LaunchTimeoutSeconds: "210",
			MaxNumRetries:        "10",
			RetryWaitTime:        "15",
			TCPNoDelay:           true,
			WorkDir:              "/var/lib/jenkins",
		},
		RetentionStrategy: computer.RetentionStrategy{
			StaplerClass: "hudson.slaves.RetentionStrategy$Always",
			Class:        "hudson.slaves.RetentionStrategy$Always",
		},
		NodeProperties: computer.NodeProperties{
			StaplerClassBag: "true",
		},
		Type: "hudson.slaves.DumbSlave",
	}

	_, err = c.CreateComputer(&req)
	if err != nil {
		log.Println("error when create computers", err)
		return
	}
}

func Test_UpdateComputer(t *testing.T) {
	config := restclient.NewConfig(baseURL, username, password)

	c, err := computer.NewForConfig(config)
	if err != nil {
		log.Println("error when create client", err)
		return
	}

	req := computer.ComputerRequest{
		Name:            "test_slave",
		NodeDescription: "",
		NumExecutors:    "4",
		RemoteFS:        "/var/lib/jenkins",
		LabelString:     "build_prod master",
		Mode:            "EXCLUSIVE", // "NORMAL"
		Launcher: computer.Launcher{
			StaplerClass:  "hudson.plugins.sshslaves.SSHLauncher",
			Class:         "hudson.plugins.sshslaves.SSHLauncher",
			Host:          "10.99.2.198",
			IncludeUser:   "false",
			CredentialsID: "slave_ssh",
			SSHHostKeyVerificationStrategy: computer.SSHHostKeyVerificationStrategy{
				StaplerClass: "hudson.plugins.sshslaves.verifiers.NonVerifyingKeyVerificationStrategy",
				Class:        "hudson.plugins.sshslaves.verifiers.NonVerifyingKeyVerificationStrategy",
			},
			Port:                 "16889",
			JavaPath:             "",
			JvmOptions:           "",
			PrefixStartSlaveCmd:  "",
			SuffixStartSlaveCmd:  "",
			LaunchTimeoutSeconds: "210",
			MaxNumRetries:        "10",
			RetryWaitTime:        "15",
			TCPNoDelay:           true,
			WorkDir:              "/var/lib/jenkins",
		},
		RetentionStrategy: computer.RetentionStrategy{
			StaplerClass: "hudson.slaves.RetentionStrategy$Always",
			Class:        "hudson.slaves.RetentionStrategy$Always",
		},
		NodeProperties: computer.NodeProperties{
			StaplerClassBag: "true",
		},
		Type: "hudson.slaves.DumbSlave",
	}

	_, err = c.UpdateComputer(&req)
	if err != nil {
		log.Println("error when update computers", err)
		return
	}
}

func Test_GetComputers(t *testing.T) {
	config := restclient.NewConfig(baseURL, username, password)

	c, err := computer.NewForConfig(config)
	if err != nil {
		log.Println("error when create client", err)
		return
	}

	var computers *computer.ComputersResponse
	computers, err = c.GetComputers()
	if err != nil {
		log.Println("error when get computers", err)
		return
	}

	log.Println(computers)
}

func Test_GetComputer(t *testing.T) {
	config := restclient.NewConfig(baseURL, username, password)

	c, err := computer.NewForConfig(config)
	if err != nil {
		log.Println("error when create client", err)
		return
	}

	var computer *computer.Computer
	computer, err = c.GetComputer("")
	if err != nil {
		log.Println("error when get computer", err)
		return
	}

	log.Println(computer)
}

func Test_DeleteComputer(t *testing.T) {
	config := restclient.NewConfig(baseURL, username, password)

	c, err := computer.NewForConfig(config)
	if err != nil {
		log.Println("error when create client", err)
		return
	}

	_, err = c.DeleteComputer("test_slave")
	if err != nil {
		log.Println("error when delete computers", err)
		return
	}
}
