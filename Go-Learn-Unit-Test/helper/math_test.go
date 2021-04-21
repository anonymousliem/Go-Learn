package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbsolute(t *testing.T) {
	t.Run("Negative test case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c, "expect 5, but actual %d", c)
	})

	t.Run("Positive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c, "expect 5, but actual %d", c)
	})
}

func TestAdd(t *testing.T) {
	t.Run("negative additional", func(t *testing.T) {
		c := Add(-1,-2)
		assert.Equal(t, -3,c)
	})


	t.Run("positive additional", func(t *testing.T) {
		c := Add(1,2)
		assert.Equal(t, 3,c)
	})
}

func TestAddTable(t *testing.T) {
	testCases := []struct {
		name   string
		a, b   int
		expect int
	}{
		{
			name: "negative and negative",
			a : -1,
			b : -1,
			expect : -2,
		},
		{
			name: "negative and positive",
			a : -1,
			b : 2,
			expect : 1,
		},
		{
			name: "positive and negative",
			a : 1,
			b : -2,
			expect : -1,
		},
		{
			name: "positif and positif",
			a : 1,
			b : 1,
			expect : 2,
		},
	}

	for _, testc := range testCases{
		t.Run(testc.name, func(t *testing.T) {
			c := Add(testc.a, testc.b)
			assert.Equal(t, testc.expect,c)
		})
	}
}