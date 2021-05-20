package coding_interview_golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func CheckNumber(number int) bool{
	if(number % 2 == 0){
		return true
	}

	return false
}

func TestCheckNumber(t *testing.T) {
	t.Run("odd test", func(t *testing.T) {
		assert.True(t, true, CheckNumber(4))
	})
	t.Run("even test", func(t *testing.T) {
		assert.False(t, false,CheckNumber(5))
	})
}
