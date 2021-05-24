package coding_interview_golang

import (
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


func TestReverseString(t *testing.T) {
	assert.Equal(t, "olah", ReverseString("halo"))
}