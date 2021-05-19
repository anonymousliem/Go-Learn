package coding_interview_golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Factorial(number int) int{
	result := 1
	if(number < 1){
		return 0
	}else {
		for i := 1; i <= number ; i++ {
			result = result * i
		}
	}
	return result
}

func FactorialRecursion(number int) int{
	if number < 1 {
		return 1
	}

	return number * FactorialRecursion(number - 1)
}

func TestFactorial(t *testing.T) {
	t.Run("with looping", func(t *testing.T) {
		result := Factorial(3)
		assert.Equal(t, 6, result)
	})

	t.Run("with recursion", func(t *testing.T) {
		result := FactorialRecursion(3)
		assert.Equal(t, 6, result)
	})
}