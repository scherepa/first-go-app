package util

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[\p{P}\s]+`)

func AlphabetPairsBytes() []byte {
	// here only one byte letters
	// Each pair is 2 letters + 1 new tab(except last)
	size := 26*3 - 1
	result := make([]byte, 0, size) // now empty with cap of size
	diff := byte('a' - 'A')
	for c := byte('A'); c <= byte('Z'); c++ {
		result = append(result, c)
		result = append(result, c+diff) // lowercase
		if c != 'Z' {
			result = append(result, '\t')
		}
	}
	return result
}

func PrintAlphabetCheckPalindroms() {
	fmt.Println(string(AlphabetPairsBytes()))
	fmt.Println("---")
	fmt.Printf("Is the word 'level' a Palindrome? - %v\n", IsPalindrome("level"))
	fmt.Printf("Is the word 'world' a Palindrome? - %v\n", IsPalindrome("world"))
	fmt.Printf("Is the word 'Level' a Palindrome? - %v\n", IsPalindrome("Level"))
	fmt.Printf("Is the following phrase 'Was it a car or a cat I saw?' a Palindrome? - %v\n", IsPalindrome("Was it a car or a cat I saw?"))
	fmt.Println("Is empty string a Palindrome? -", IsPalindrome(""))
	fmt.Printf("Is the word 'אבא' a Palindrome? - %v\n", IsPalindrome("אבא"))
}

func IsPalindrome(s string) bool {
	// lets remove spaces and punctuation
	cleanedString := re.ReplaceAllString(strings.ToLower(s), "")
	// important for unicode(Hebrew, Chinese, emojis, etc.)
	// for which each letter can contain several blocks in it(multiple bytes)
	runeFromString := []rune(cleanedString)
	length := len(runeFromString)
	if length <= 1 {
		return true
	}
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		if runeFromString[i] != runeFromString[j] {
			return false
		}
	}
	return true
}
