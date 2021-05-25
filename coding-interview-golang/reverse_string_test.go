package coding_interview_golang

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ReverseString(word string) (result string){

	if len(word) == 1 {
		return  word
	}

		for _, i := range word {
			result = string(i) + result
			//fmt.Println(result)
		}

		return result

}

func reverseWithRecursive(word string) string{
	if len(word) == 1{
		return word
	}
	fmt.Println(reverseWithRecursive(word[1:]) + word[0:1])
	return reverseWithRecursive(word[1:]) + word[0:1]
}

func reverse(str string) {
	if len(str) == 1 {
		fmt.Print(str)
		return
	}
	reverse(str[1:])
	fmt.Print(str[0:1])
}

func TestReverseString(t *testing.T) {
	t.Run("with loop", func(t *testing.T) {
		assert.Equal(t, "olah", ReverseString("halo"))
	})

	t.Run("with recursive", func(t *testing.T) {
		assert.Equal(t, "olah", reverseWithRecursive("halo"))
	})


}