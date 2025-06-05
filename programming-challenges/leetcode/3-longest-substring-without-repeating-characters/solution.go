package main

func main() {
	result := lengthOfLongestSubstring(" ")
	println(result)
}

func lengthOfLongestSubstring(s string) int {
	var i, j, maxlenght int
	i = 0
	j = 0
	maxlenght = 0
	letters := []rune(s)
	runemap := make(map[rune]int, len(s))

	for j < len(letters) {
		nextLetter := letters[j]
		ocurrencyIndex, exists := runemap[nextLetter]
		if exists && ocurrencyIndex >= i {
			// Check if the current window is the longest
			if lenght := j - i; lenght > maxlenght {
				maxlenght = lenght
			}
			// jump to the next possible valid window
			i = ocurrencyIndex + 1
		}
		runemap[nextLetter] = j
		j++
	}

	// Last check if the maximum was the last window
	if lenght := j - i; lenght > maxlenght {
		maxlenght = lenght
	}
	return maxlenght
}
