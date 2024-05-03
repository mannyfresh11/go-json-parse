package main

import (
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

func (l *Logs) UnmarshalJSON(b []byte) error {
	var fields []string
	err := json.Unmarshal(b, &fields)
	if err == nil {
		*l = Logs{}
		fields = append(fields, l.RemoteAddr)
		fields = append(fields, l.RemoteUser)
		fields = append(fields, l.Time)
		fields = append(fields, l.Request)
		fields = append(fields, l.Status)
		fields = append(fields, l.BodyByte)
		fields = append(fields, l.RequestTime)
		fields = append(fields, l.HTTPRef)
		fields = append(fields, l.HTTPUserAgent)
		return nil
	}
	return nil
}

func main() {
	//jsonFile, err := os.ReadFile("./jsonexample.log")
	jsonFile, err := os.ReadFile("./jsonexample.log")
	check(err)
	//fmt.Printf("this is what is inside jsonFile: %v\n", jsonFile)

	// defer jsonFile.Close()

	var logs []*Logs
	// data := make([]byte, 1024)
	// _, err = jsonFile.Read(data)
	// check(err)

	//os.Stdout.Write(data)
	err = json.Unmarshal(jsonFile, &logs)
	check(err)

	for _, log := range logs {
		// fmt.Printf("this is what is inside logs: %v\n", logs[:count])
		fmt.Printf("Request Address: %v\n", log.RemoteAddr)
		fmt.Printf("Request User: %v\n", log.RemoteUser)
		fmt.Printf("Time: %v\n", log.Time)
		fmt.Printf("Requet: %v\n", log.Request)
		fmt.Printf("Status: %v\n", log.Status)
		fmt.Printf("Body byte sent: %v\n", log.BodyByte)
		fmt.Printf("Request Time: %v\n", log.RequestTime)
		fmt.Printf("Http Referrer: %v\n", log.HTTPRef)
		fmt.Printf("Http User Agent: %v\n", log.HTTPUserAgent)
	}
}
