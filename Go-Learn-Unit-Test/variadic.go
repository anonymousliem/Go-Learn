package main

import (
	"fmt"
)

func sum(nums ...int) int {
	s := 0
	for _, i := range nums {
		s += i
	}
	return s
}


func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))
}