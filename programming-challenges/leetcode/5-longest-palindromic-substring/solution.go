package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("cbbd"))

}

func longestPalindrome(s string) string {
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
