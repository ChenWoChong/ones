package deploy

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh"
)

func main() {
	// SSH client configuration
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server
	conn, err := ssh.Dial("tcp", "remote-server:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer conn.Close()

	// Create a new session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Check if Docker is installed
	dockerInstalled, err := checkIfDockerInstalled(session)
	if err != nil {
		log.Fatalf("Failed to check if Docker is installed: %s", err)
	}

	// Install Docker if it is not installed
	if !dockerInstalled {
		err = installDocker(session)
		if err != nil {
			log.Fatalf("Failed to install Docker: %s", err)
		}
	}

	// Check if Kubernetes is installed
	k8sInstalled, err := checkIfK8sInstalled(session)
	if err != nil {
		log.Fatalf("Failed to check if Kubernetes is installed: %s", err)
	}

	// Install Kubernetes if it is not installed
	if !k8sInstalled {
		err = installK8s(session)
		if err != nil {
			log.Fatalf("Failed to install Kubernetes: %s", err)
		}
	}

	// Add the worker node to the Kubernetes cluster
	err = addWorkerNode(session)
	if err != nil {
		log.Fatalf("Failed to add worker node: %s", err)
	}
}

// Check if Docker is installed
func checkIfDockerInstalled(session *ssh.Session) (bool, error) {
	cmd := "which docker"
	output, err := runCommand(session, cmd)
	if err != nil {
		return false, err
	}
	if output == "" {
		return false, nil
	}
	return true, nil
}

// Install Docker
func installDocker(session *ssh.Session) error {
	cmd := "sudo apt-get update && apt-get install -y docker-ce docker-ce-cli containerd.io"
	output, err := runCommand(session, cmd)
	if err != nil {
		return  err
	}
	if output == "" {
		return nil
	}
	return nil
}

func main() {
	// SSH client configuration
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server
	conn, err := ssh.Dial("tcp", "remote-server:22", config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer conn.Close()

	// Create a new session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	// Check if Docker is installed
	dockerInstalled, err := checkIfDockerInstalled(session)
	if err != nil {
		log.Fatalf("Failed to check if Docker is installed: %s", err)
	}

	// Install Docker if it is not installed
	if !dockerInstalled {
		err = installDocker(session)
		if err != nil {
			log.Fatalf("Failed to install Docker: %s", err)
		}
	}

	// Check if Kubernetes is installed
	k8sInstalled, err := checkIfK8sInstalled(session)
	if err != nil {
		log.Fatalf("Failed to check if Kubernetes is installed: %s", err)
	}

	// Install Kubernetes if it is not installed
	if !k8sInstalled {
		err = installK8s(session)
		if err != nil {
			log.Fatalf("Failed to install Kubernetes: %s", err)
		}
	}

	// Add the worker node to the Kubernetes cluster
	err = addWorkerNode(session)
	if err != nil {
		log.Fatalf("Failed to add worker node: %s", err)
	}
}

// Check if Docker is installed
func checkIfDockerInstalled(session *ssh.Session) (bool, error) {
	cmd := "which docker"
	output, err := runCommand(session, cmd)
	if err != nil {
		return false, err
	}
	if		return true, nil
	}
	return true, nil
}

// Install Kubernetes
func installK8s(session *ssh.Session) error {
	cmds := []string{
		"sudo apt-get update",
		"sudo apt-get install -y apt-transport-https ca-certificates curl",
		"curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg",
		"echo \"deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main\" | sudo tee /etc/apt/sources.list.d/kubernetes.list",
		"sudo apt-get update",
		"sudo apt-get install -y kubelet kubeadm kubectl",
		"sudo apt-mark hold kubelet kubeadm kubectl",
	}
	for _, cmd := range cmds {
		err := runCommandWithOutput(session, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Add the worker node to the Kubernetes cluster
func addWorkerNode(session *ssh.Session) error {
	// Get the join command from the master node
	cmd := "sudo kubeadm token create --print-join-command"
	output, err := runCommand(session, cmd)
	if err != nil {
		return err
	}

	// Join the worker node to the cluster
	cmd = strings.TrimSpace(output)
	err = runCommandWithOutput(session, cmd)
	if err != nil {
		return err
	}

	return nil
}

// Run a command on the remote server and return the output
func runCommand(session *ssh.Session, cmd string) (string, error) {
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", fmt.Errorf("command %q failed: %s", cmd, output)
	}
	return strings.TrimSpace(string(output)), nil
}

// Run a command on the remote server and print the output to stdout
func runCommandWithOutput(session *ssh.Session, cmd string) error {
	fmt.Printf("Running command: %s\n", cmd)
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return fmt.Errorf("command %q failed: %s", cmd, output)
	}
	fmt.Printf("%s\n", output)
	return nil
}

