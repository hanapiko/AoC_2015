package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	secretKey := "iwrupvqb"
	result := mineAdventCoins(secretKey)
	res := mineAdventCoins6(secretKey)
	fmt.Println(result)
	fmt.Println(res)
}

func mineAdventCoins(secretKey string) int {
	number := 0
	for {
		// Concatenate the secret key with the current number
		input := fmt.Sprintf("%s%d", secretKey, number)
		// Compute the MD5 hash
		hash := md5.Sum([]byte(input))
		hashString := hex.EncodeToString(hash[:])
		// Check if the hash starts with five zeroes
		if hashString[:5] == "00000" {
			return number
		}
		// Increment the number for the next iteration
		number++
	}
}

func mineAdventCoins6(secretKey string) int {
	number := 0
	for {
		input := fmt.Sprintf("%s%d", secretKey, number)
		hash := md5.Sum([]byte(input))
		hashString := hex.EncodeToString(hash[:])
		if hashString[:6] == "000000" {
			return number
		}
		number++
	}
}
