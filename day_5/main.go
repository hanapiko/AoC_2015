package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the data file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var count1 int
	var count2 int
	scanner := bufio.NewScanner(file)
	//slc := []string{}
	// Read through the file line by line and count vowels
	for scanner.Scan() {
		line := scanner.Text()
		if IsNice(line) {
			count1++
		}
		if isNicePartTwo(line) {
			count2++
		}
		 
	}
	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	//fmt.Println(countNice((slc)))
	fmt.Println(count1)
	fmt.Println(count2)
}

func IsVowel(s string) bool {
	vowels := "aeiou"
	count := 0
	for i := 0; i < len(s); i++ {
		if strings.ContainsRune(vowels, rune(s[i])) {
			count++
		}
	}
	return count >= 3
}

func double(s string) bool {
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func forbidden(s string) bool {
	forbiddenSubstrings := []string{"ab", "cd", "pq", "xy"}
	for i := 0; i < len(s)-1; i++ {
		for _, sub := range forbiddenSubstrings {
			if strings.Contains(s, sub) {
				return false
			}
		}
	}

	return true
}

func IsNice(s string) bool {
	if forbidden(s) && IsVowel(s) && double(s) {
		return true
	}
	return false
}

// part two
func isNicePartTwo(s string) bool {
    return hasRepeatedPair(s) && hasRepeatingWithOneBetween(s)
}

func hasRepeatedPair(s string) bool {
    seen := make(map[string]int)

    for i := 0; i < len(s)-1; i++ {
        pair := s[i : i+2]
        seen[pair]++
        if seen[pair] == 2 {
            return true
        }
    }

    return false
}

func hasRepeatingWithOneBetween(s string) bool {
    for i := 0; i < len(s)-2; i++ {
        if s[i] == s[i+2] {
            return true
        }
    }
    return false
}