package coding_interview_golang

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func PalindromeWithTemp(word string) bool {
	temp := ""
	for i := len(word) -1 ; i >= 0; i-- {
		temp = temp + string(word[i])
	}

	return strings.EqualFold(temp,word)
}

func PalindromeCompareChar (word string) bool{
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word) - i -1] {
			return false
		}
	}
	return true
}

func PalindromeWithRecursive (word string, i int) bool{
	if i < len(word)/2 {
		if word[i] != word[len(word) - i -1] {
			return false
		}else{
			return PalindromeWithRecursive(word, i + 1)
		}
	}else{
		return true
	}
}

func TestPalindrome(t *testing.T) {
	PalindromeWithTemp("kodok")
	assert.True(t, PalindromeWithTemp("a"))
	assert.True(t, PalindromeWithTemp("aba"))
	assert.True(t, PalindromeWithTemp("kodok"))
	assert.True(t, PalindromeWithTemp(""))

	assert.False(t, PalindromeWithTemp("ab"))
	assert.False(t, PalindromeWithTemp("abad"))
	assert.False(t, PalindromeWithTemp("kodcok"))
	assert.False(t, PalindromeWithTemp("awokowwok"))
}

func TestPalindromeCompare(t *testing.T) {
	PalindromeCompareChar("kodok")
	assert.True(t, PalindromeCompareChar("a"))
	assert.True(t, PalindromeCompareChar("aba"))
	assert.True(t, PalindromeCompareChar("kodok"))
	assert.True(t, PalindromeCompareChar(""))

	assert.False(t, PalindromeCompareChar("ab"))
	assert.False(t, PalindromeCompareChar("abad"))
	assert.False(t, PalindromeCompareChar("kodcok"))
	assert.False(t, PalindromeCompareChar("awokowwok"))
}

func TestPalindromeWithRecursive(t *testing.T) {
	PalindromeCompareChar("kodok")
	assert.True(t, PalindromeWithRecursive("a",0))
	assert.True(t, PalindromeWithRecursive("aba",0))
	assert.True(t, PalindromeWithRecursive("kodok",0))
	assert.True(t, PalindromeWithRecursive("",0))

	assert.False(t, PalindromeWithRecursive("ab",0))
	assert.False(t, PalindromeWithRecursive("abad",0))
	assert.False(t, PalindromeWithRecursive("kodcok",0))
	assert.False(t, PalindromeWithRecursive("awokowwok",0))
}