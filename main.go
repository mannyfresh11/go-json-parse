package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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

	// conn := NewSSHConnection("administrator", "servicerequest.dfci.harvard.edu:22")

	// defer conn.Close()

	// sess, err := NewSession("cat /var/log/nginx/access.log", conn)
	// if err != nil {
	// 	log.Fatalf("Error from session is: %v\n", err)
	// }
	// defer sess.Close()

	jsonFile, err := os.ReadFile("./json.log")
	check(err)

	buf := bytes.NewBuffer(jsonFile)
	dec := json.NewDecoder(buf)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Remote Adress", "Remote User", "Date and Time", "Request", "Status", "Body Byte", "Request Time", "HTTP Ref", "HTTP User Agent"}, rowConfigAutoMerge)
	t.SetAutoIndex(true)
	// t.SetColumnConfigs([]table.ColumnConfig{
	// 	{Number: 1, AutoMerge: true},
	// 	{Number: 2, AutoMerge: true},
	// 	{Number: 3, AutoMerge: true},
	// 	{Number: 4, AutoMerge: true},
	// 	{Number: 5, AutoMerge: true},
	// 	{Number: 6, AutoMerge: true},
	// 	{Number: 7, AutoMerge: true},
	// })
	t.SetStyle(table.StyleLight)
	t.SetAllowedRowLength(250)
	t.Style().Options.SeparateRows = true

	for dec.More() {
		var logs Logs

		err := dec.Decode(&logs)
		check(err)

		t.AppendRow(table.Row{logs.RemoteAddr, logs.RemoteUser, logs.Time, logs.Request, logs.Status, logs.BodyByte, logs.RequestTime, logs.HTTPRef, logs.HTTPUserAgent}, rowConfigAutoMerge)
		// fmt.Printf("Request Address: %v\n", logs.RemoteAddr)
		// fmt.Printf("Request User: %v\n", logs.RemoteUser)
		// fmt.Printf("Time: %v\n", logs.Time)
		// fmt.Printf("Requet: %v\n", logs.Request)
		// fmt.Printf("Status: %v\n", logs.Status)
		// fmt.Printf("Body byte sent: %v\n", logs.BodyByte)
		// fmt.Printf("Request Time: %v\n", logs.RequestTime)
		// fmt.Printf("Http Referrer: %v\n", logs.HTTPRef)
		// fmt.Printf("Http User Agent: %v\n\n", logs.HTTPUserAgent)
	}
	fmt.Println(t.Render())
}
