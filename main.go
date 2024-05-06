package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	conn := SshConnecting("manny", "192.168.1.101:22")

	session, err := conn.NewSession()
	check(err)
	var b bytes.Buffer

	session.Stdout = &b

	err = session.Run("/usr/bin/whoami")
	check(err)

	fmt.Println(b.String())

	jsonFile, err := os.ReadFile("./jsonexample.log")
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
