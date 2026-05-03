package main

import (
	"fmt"
)

func main() {
	rooms := map[string]float64{
		"room 1": 1.00,
		"room 2": 2.00,
		"room 3": 3.00,
		"room 4": 4.00,
	}

	fmt.Println(rooms)
	fmt.Println(rooms["room 3"])

	// Looping through a map
	for key, value := range rooms {
		fmt.Println(key, "-", value)
	}
}
