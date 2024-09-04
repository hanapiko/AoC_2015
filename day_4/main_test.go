package main

import (
    "testing"
)

// TestMineAdventCoins tests the mineAdventCoins function.
func TestMineAdventCoins(t *testing.T) {
    secretKey := "iwrupvqb"
    expected := 346386
    result := mineAdventCoins(secretKey)

    if result != expected {
        t.Errorf("mineAdventCoins(%s) = %d; want %d", secretKey, result, expected)
    }
}

// TestMineAdventCoins6 tests the mineAdventCoins6 function.
func TestMineAdventCoins6(t *testing.T) {
    secretKey := "iwrupvqb"
    expected := 9958218
    result := mineAdventCoins6(secretKey)

    if result != expected {
        t.Errorf("mineAdventCoins6(%s) = %d; want %d", secretKey, result, expected)
    }
}