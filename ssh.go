package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func NewSSHConnection(user, host string) *ssh.Client {

	key, err := os.ReadFile("/home/manny/.ssh/id_ed25519")
	if err != nil {
		fmt.Printf("Error reading key: %v\n", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
	}

	hostKeyCallback, err := knownhosts.New("/home/manny/.ssh/known_hosts")
	if err != nil {
		fmt.Printf("Error geting known host: %v\n", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback:   hostKeyCallback,
		HostKeyAlgorithms: []string{ssh.KeyAlgoED25519},
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatalf("Error connecting to client: %v\n", err)
	}

	return client
}

func NewSession(cmd string, conn *ssh.Client) (*ssh.Session, error) {

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Error occured creating session: %v\n", err)
	}

	var b bytes.Buffer

	session.Stdout = &b

	err = session.Run(cmd)
	if err != nil {
		log.Fatalf("Error running comand: %v\n", err)
	}

	os.WriteFile("./json.log", b.Bytes(), 0666)

	return session, err
}
