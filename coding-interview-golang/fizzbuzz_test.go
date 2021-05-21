package coding_interview_golang

import (
	"fmt"
	"testing"
)

func FizzBuzz(number int)  {
	for i := 1; i <= number; i++ {
		if(i % 3 == 0 && i % 5 ==0){
			fmt.Println("FizzBuzz")
		}else if(i % 3 == 0) {
			fmt.Println("Fizz")
		}else if (i % 5 == 0) {
			fmt.Println("buzz")
		}else{
			fmt.Println(i)
		}
	}
}

func TestFizzbuzz(t *testing.T) {
	FizzBuzz(100)
}