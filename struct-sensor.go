package main

import (
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID   string
	Time time.Time
}

type DoorEvent struct {
	Event
	Action string //open, close
}

type TemperatureEvent struct {
	Event
	Value float32
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}

	evt := DoorEvent{
		Event:  Event{id, time},
		Action: action,
	}

	return &evt, nil
}

func main() {
	evt, err := NewDoorEvent("fron door", time.Now(), "open")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%+v\n", evt)
}
