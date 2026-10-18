package main

import ("fmt";"sort")

type Event struct {
	At int
	Name string
}

func main() {
	events := []Event {
		{At: 50, Name: "message arrives"}, {At: 35, Name: "time fires"}, {At: 20, Name: "node crashes"},
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].At < events[j].At
	})
	fmt.Println(events)
}