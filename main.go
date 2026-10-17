package main

import "fmt"

type Event struct {
	At int
	Name string
}

func main() {
	e := Event{At: 20, Name: "timer fires"}
	fmt.Println(e.At, e.Name)
}