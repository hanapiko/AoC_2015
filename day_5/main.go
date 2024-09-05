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
	scanner := bufio.NewScanner(file)
	slc := []string{}
	// Read through the file line by line and count vowels
	for scanner.Scan() {
		line := scanner.Text()

		slc = append(slc, line)
	}
	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(Nice(slc))
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
		for i := 0; i < len(s)-1; i++{
			for _, sub := range forbiddenSubstrings {
				if strings.Contains(s, sub) {
					return false
				}
			}
	}

	// Return true if all conditions are met
	// return Count >= 3 && hasDoubleLetter
	return true
}

func IsNice(s string)bool{
	if forbidden(s) && IsVowel(s) && double(s){
			return true
		}
	return false
}

func Nice(str []string) int {
	count := 0

	for _, s:= range str{
		if IsNice(s){
			count++
		}
	}
	return count
}
