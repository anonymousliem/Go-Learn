package coding_interview_golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func LinearSearch(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return true
		}
	}
	return false
}

func TestLinear(t *testing.T) {
	arr := []int{4, 2, 5, 8, 3, 9, 7}
	assert.False(t,LinearSearch(arr, 6) )
	assert.True(t,LinearSearch(arr, 2) )
}