package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	username = os.Getenv("REMOTE_USER")
	hostname = os.Getenv("REMOTE_HOST")
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

	conn := NewSSHConnection(username, hostname)

	defer conn.Close()

	err := GetLogFile("cat /var/log/nginx/access.log", conn)
	if err != nil {
		log.Fatalf("Error from session is: %v\n", err)
	}

	jsonFile, err := os.ReadFile("./json.log")
	check(err)

	buf := bytes.NewBuffer(jsonFile)
	dec := json.NewDecoder(buf)

	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Remote Adress", "Remote User", "Date and Time", "Request", "Status", "Body Byte", "Request Time"}, rowConfigAutoMerge)
	t.SetAutoIndex(true)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
		// 	{Number: 2, AutoMerge: true},
		{Number: 3, AutoMerge: true},
		// 	{Number: 4, AutoMerge: true},
		{Number: 5, AutoMerge: true},
		// 	{Number: 6, AutoMerge: true},
		//{Number: 7, AutoMerge: true},
		//{Number: 8, AutoMerge: true},
		//{Number: 9, AutoMerge: true},
	})
	t.SetStyle(table.StyleLight)
	t.SetAllowedRowLength(100)
	t.Style().Options.SeparateRows = true

	fmt.Println("Table creation started.")
	for dec.More() {
		var logs Logs

		err := dec.Decode(&logs)
		check(err)

		localTimeZN, err := time.LoadLocation("America/New_York")
		check(err)

		layout := "2006-01-02T15:04:05+00:00"

		localTime, err := time.ParseInLocation(layout, logs.Time, localTimeZN)
		check(err)

		t.AppendRow(table.Row{logs.RemoteAddr, logs.RemoteUser, localTime, logs.Request, logs.Status, logs.BodyByte, logs.RequestTime}, rowConfigAutoMerge)
	}

	os.WriteFile("./jsontable.log", []byte(t.Render()), 0666)
	fmt.Println("Table creation finished.")
}
