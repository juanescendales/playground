package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("cbbd"))

}

func longestPalindrome(s string) string {
	var oddPalindrome, evenPalindrome string
	maxPalindrome, size := s[:1], len(s)
	for k := range size - 1 {
		oddPalindrome = expandPalindrome(s, k, k)
		if len(oddPalindrome) > len(maxPalindrome) {
			maxPalindrome = oddPalindrome
		}
		evenPalindrome = expandPalindrome(s, k, k+1)
		if len(evenPalindrome) > len(maxPalindrome) {
			maxPalindrome = evenPalindrome
		}
	}
	return maxPalindrome
}

func expandPalindrome(s string, c1, c2 int) string {
	l, r := c1, c2
	size := len(s)
	palindrome, palindromeLen := "", 0

	for (l >= 0 && r < size) && (s[l] == s[r]) {
		newLength := r - l + 1
		if newLength > palindromeLen {
			palindrome = s[l : r+1]
			palindromeLen = newLength
		}
		l--
		r++
	}
	return palindrome
}

func LongestPalindromeO3Solution(s string) string {
	var j, size int
	letters := []rune(s)
	size = len(letters)
	j = size

	for j > 1 {
		start := 0
		end := j
		for end <= size {
			substring := letters[start:end]
			if isPalindrome(substring) {
				return string(substring)
			}
			end++
			start++
		}
		j--
	}
	return string(letters[0])
}

func isPalindrome(word []rune) bool {
	var i, j int
	size := len(word)
	i = 0
	j = size - 1

	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}
