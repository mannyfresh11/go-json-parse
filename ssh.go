package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func SshConnecting(user, host string) *ssh.Client {

	var hostKey ssh.PublicKey

	key, err := os.ReadFile("~/.ssh/id_rsa")
	if err != nil {
		fmt.Printf("Error reading key: %v\n", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
	}
	defer client.Close()

	return client
}
