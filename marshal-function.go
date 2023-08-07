package main

import (
	"fmt"
	"log"
)

type ClickEvent struct {
}

type HoverEvent struct {
}

var eventCounts = make(map[string]int)

func recordEvent(evt interface{}) {
	switch evt.(type) {
	case *ClickEvent:
		eventCounts["click"]++
	case *HoverEvent:
		eventCounts["hover"]++
	default:
		log.Printf("Warning: unknown event: %#v of type %T\n", evt, evt)
	}
}

func main() {
	recordEvent(&ClickEvent{})
	recordEvent(&HoverEvent{})
	recordEvent(&ClickEvent{})
	recordEvent(3)
	fmt.Println("event counts", eventCounts)
}
