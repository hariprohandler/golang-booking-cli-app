package main

import (
	"fmt"
)

func getConsecutiveSubstrings(s string) []string {
	var substrings []string

	for index := 0; index < len(s); index++ {
		currentSubstring := string(s[index])
		substrings = append(substrings, currentSubstring)

		for i := index; i < len(s)-1; i++ {
			if s[i+1]-s[i] == 1 {
				currentSubstring += string(s[i+1])
				substrings = append(substrings, currentSubstring)
			} else {
				break
			}
		}
	}
	return substrings
}

func main() {
	var str1, str2 string
	fmt.Println("Enter first string:")
	fmt.Scan(&str1)
	fmt.Println("Enter second string:")
	fmt.Scan(&str2)

	substrings1 := getConsecutiveSubstrings(str1)
	substrings2 := getConsecutiveSubstrings(str2)

	// Store substrings2 in a map for fast lookup
	subMap := make(map[string]bool)
	for _, s := range substrings2 {
		subMap[s] = true
	}

	// Find common substrings
	fmt.Println("\nCommon consecutive substrings:")
	found := false
	for _, s := range substrings1 {
		if subMap[s] {
			fmt.Println(s)
			found = true
		}
	}
	if !found {
		fmt.Println("No common substrings found.")
	}
}
