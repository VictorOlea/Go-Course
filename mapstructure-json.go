package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// StartJob is a request to start a job
type StartJob struct {
	Type  string
	User  string
	Count int
}

// JobStatus is a request for job status
type JobStatus struct {
	Type string
	ID   string
}

func handleStart(req StartJob) error {
	fmt.Printf("Start: %#v\n", req)
	return nil
}

func handleStatus(req JobStatus) error {
	fmt.Println("Status: %#v\n", req)
	return nil
}

func handleRequest(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	val, ok := m["type"]
	if !ok {
		return fmt.Errorf("type 'type' missing from JSON")
	}

	typ, ok := val.(string)
	if !ok {
		return fmt.Errorf(" 'type' is not a string")
	}

	switch typ {
	case "start":
		var sj StartJob
		if err := mapstructure.Decode(m, &sj); err != nil {
			return fmt.Errorf("bad 'start' request: %W", err)
		}
		return handleStart(sj)

	case "status":
		var js JobStatus
		if err := mapstructure.Decode(m, &js); err != nil {
			return fmt.Errorf("Bad 'status' request: %w", err)
		}
		return handleStatus(js)
	}
	return fmt.Errorf("Unknown Request Type: %q", typ)
}

func main() {
	data := []byte(`{"type": "start", "user": "joe", "count": 7}`)
	if err := handleRequest(data); err != nil {
		fmt.Println("ERROR: ", err)
	}

	data = []byte(`{"type": "status", "id": "seven"}`)
	if err := handleRequest(data); err != nil {
		fmt.Println("ERROR: ", err)
	}
}
