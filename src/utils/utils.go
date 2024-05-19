package utils

import "fmt"

func PrintSeprater() {
	fmt.Printf("-----------------------------\n")
}

type alterCounter int

func New(value int) alterCounter {
	return alterCounter(value)
}

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Right int
}
