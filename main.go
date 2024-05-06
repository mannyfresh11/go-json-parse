package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Logs struct {
	Time          string `json:"time_local"`
	RemoteAddr    string `json:"remote_addr"`
	RemoteUser    string `json:"remote_user"`
	Request       string `json:"request"`
	Status        string `json:"status"`
	BodyByte      string `json:"body_bytes_sent"`
	RequestTime   string `json:"request_time"`
	HTTPRef       string `json:"http_referrer"`
	HTTPUserAgent string `json:"http_user_agent"`
}

func check(e error) {
	if e != nil {
		fmt.Printf("error found: %v\n", e)
	}
}

func main() {

	conn := NewSSHConnection("administrator", "servicerequest.dfci.harvard.edu:22")

	defer conn.Close()

	sess, err := NewSession("cat /var/log/nginx/access.log", conn)
	if err != nil {
		log.Fatalf("Error from session is: %v\n", err)
	}
	defer sess.Close()

	jsonFile, err := os.ReadFile("./json.log")
	check(err)

	buf := bytes.NewBuffer(jsonFile)
	dec := json.NewDecoder(buf)

	for dec.More() {
		var logs Logs

		err := dec.Decode(&logs)
		check(err)

		fmt.Printf("Request Address: %v\n", logs.RemoteAddr)
		fmt.Printf("Request User: %v\n", logs.RemoteUser)
		fmt.Printf("Time: %v\n", logs.Time)
		fmt.Printf("Requet: %v\n", logs.Request)
		fmt.Printf("Status: %v\n", logs.Status)
		fmt.Printf("Body byte sent: %v\n", logs.BodyByte)
		fmt.Printf("Request Time: %v\n", logs.RequestTime)
		fmt.Printf("Http Referrer: %v\n", logs.HTTPRef)
		fmt.Printf("Http User Agent: %v\n\n", logs.HTTPUserAgent)
	}
}
