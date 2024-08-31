package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	houses := House(string(file))
	new := SantaRobo(string(file))
	fmt.Println(houses)
	fmt.Println(new)
}

func House(s string) int {
	// coordinates and map to track visited houses
	x, y := 0, 0
	seen := make(map[string]bool)

	// Mark the starting house as visited
	seen[fmt.Sprintf("%d,%d", x, y)] = true

	for _, ch := range s {
		switch ch {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		// Mark current house visited
		seen[fmt.Sprintf("%d,%d", x, y)] = true
	}

	// count of unique houses visited
	return len(seen)
}

func SantaRobo(s string) int {
	// coordinates and map to track visited houses
	sx, sy := 0, 0
	rx, ry := 0, 0

	seen := make(map[string]bool)

	// Mark the starting house as visited
	seen[fmt.Sprintf("%d,%d", sx, sy)] = true

	for i, ch := range s {
		if i%2 == 1 {
			switch ch {
			case '^':
				sy++
			case 'v':
				sy--
			case '>':
				sx++
			case '<':
				sx--
			}
			// Mark current house visited
			seen[fmt.Sprintf("%d,%d", sx, sy)] = true
		} else {
			switch ch {
			case '^':
				ry++
			case 'v':
				ry--
			case '>':
				rx++
			case '<':
				rx--
			}
			// Mark current house visited
			seen[fmt.Sprintf("%d,%d", rx, ry)] = true
		}
	}

	// count of unique houses visited by both
	return len(seen)
}
