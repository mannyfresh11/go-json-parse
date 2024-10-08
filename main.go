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
	t.AppendHeader(table.Row{
		"Remote Adress",
		"Date and Time",
		"Request",
		"Status",
		"Body Byte"},
		rowConfigAutoMerge)
	t.SetAutoIndex(true)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 3, AutoMerge: true, WidthMax: 48},
		{Number: 5, AutoMerge: true},
	})
	t.SetStyle(table.StyleLight)
	t.SetAllowedRowLength(200)
	t.Style().Options.SeparateRows = true

	layout := "2006-01-02T15:04:05+00:00"
	localTimeZN, err := time.LoadLocation("America/New_York")
	check(err)

	var localTime time.Time

	fmt.Println("Table creation started...")
	for dec.More() {
		var logs Logs

		err := dec.Decode(&logs)
		check(err)

		parseTime, err := time.Parse(layout, logs.Time)
		check(err)

		localTime = parseTime.In(localTimeZN)
		formattedTime := localTime.Format("2006-01-02 15:04:05")

		t.AppendRow(table.Row{
			logs.RemoteAddr,
			formattedTime,
			logs.Request,
			logs.Status,
			logs.BodyByte},
			rowConfigAutoMerge)
	}

	os.WriteFile("./jsontable.log", []byte(t.Render()), 0666)
	fmt.Println("Table creation finished.")
}
