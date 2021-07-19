package main

import "fmt"

type Cities struct {
	name string
	location [2]int
}

func main() {

	// Create empty slice of struct pointers.
	cities := []*Cities{}

	// Create struct and append it to the slice.
	ct := new(Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 0
	cities = append(cities, ct)

	// Create another struct.
	ct = new(Cities)
	ct.name = "Sydney"
	ct.location[0] = 34
	ct.location[1] = 51
	cities = append(cities, ct)

	for i := range(cities) {
		c := cities[i]
		fmt.Println("City:", *c)
	}
}