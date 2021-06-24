package coding_interview_golang

import (
	"fmt"
	"math"
	"testing"
)

func printNumberPrime(num1, num2 int) int {
	if num1<2 || num2<2{
		fmt.Println("Numbers must be greater than 2.")
		return 0
	}
	for num1 <= num2 {
		isPrime := true
		for i:=2; i<=int(math.Sqrt(float64(num1))); i++{
			if num1 % i == 0{
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		 num1++
	}
	fmt.Println()
	return 0

}

func TestPrintPrimeNumber(t *testing.T) {
	printNumberPrime(5,19)
}