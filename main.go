package main

import (
	"fmt"
	"sort"
)

type Event struct {
	At   int
	Name string
	Seq  int
}

func main() {
	events := []Event{
		{At: 50, Name: "message arrives", Seq: 1}, {At: 35, Name: "time fires", Seq: 3}, {At: 20, Name: "node crashes", Seq: 4}, {At: 35, Name: "step 2", Seq: 2},
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].At == events[j].At {
			return events[i].Seq < events[j].Seq
		}
		return events[i].At < events[j].At
	})
	fmt.Println(events)
}
