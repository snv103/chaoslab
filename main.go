package main

import (
	"fmt"
	"container/heap"
)

type Event struct {
	At   int
	Name string
	Seq  int
}

type EventHeap []Event
func(h EventHeap) Len()int {return len(h)}
func(h EventHeap) Less(i, j int)bool { 
	if h[i].At == h[j].At {
		return h[i].Seq < h[j].Seq
	}
	return h[i].At < h[j].At
}
func(h EventHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func(h *EventHeap) Push (x any) {
	*h = append(*h, x.(Event))
}
func(h *EventHeap) Pop() any {
	old := *h
	n := len(*h)
	x := old[n - 1]
	*h = old[: n - 1]
	return x
}

func main() {
	events := &EventHeap {
		{At: 20, Name: "time fires", Seq: 1}, {At: 60, Name: "end", Seq: 4}, {At: 35, Name: "step 3", Seq: 3}, {At: 35, Name: "node crashes", Seq: 2}, {At: 60, Name: "step 4", Seq: 6}, {At: 25, Name: "step 2", Seq: 5}, {At: 10, Name: "message arrives", Seq: 7},
	}
	heap.Init(events)
	heap.Push(events, Event{At: 20, Name: "new event", Seq: 8})
	for events.Len() > 0 {
		e:= heap.Pop(events).(Event)
		fmt.Println(e.At, e.Name, e.Seq)
	}
}
