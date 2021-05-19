package coding_interview_golang

import (
	"fmt"
	"strconv"
	"testing"
)

func FibonacciRecursion(number int) int{
	if number <= 1{
		return  number
	}

	result := FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
	return  result
}

func FibonacciLoop(number int) int{
	f:= make([]int, number+1, number+2)

	if number < 2 {
		f = f[0:2]
	}

	f[0] = 0
	f[1] = 1

	for i := 2; i <= number ; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[number]
}

func TestFibonacci(t *testing.T) {
	t.Run("fibonacci with recursion", func(t *testing.T) {
		for i := 0; i <= 9; i++ {
			fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
		}
	})

	t.Run("fibonacci with loop", func(t *testing.T) {
		for i := 0; i <= 5; i++ {
			fmt.Print(strconv.Itoa(FibonacciLoop(i)) + " ")
		}
	})
}