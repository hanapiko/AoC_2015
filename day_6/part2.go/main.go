package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 1000

func main() {
	// Initialize the grid with zeros (all lights start with zero brightness)
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	// Open the file
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read and process each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Check if the line is not empty
			err := Instruction(grid, line)
			if err != nil {
				fmt.Println("Error processing instruction:", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate the total brightness
	totalBrightness := calculateTotalBrightness(grid)
	fmt.Printf("Total brightness: %d\n", totalBrightness)
}

// Processes a single instruction and updates the grid
func Instruction(grid [][]int, instruction string) error {
	parts := strings.Fields(instruction)
	if len(parts) < 4 {
		return fmt.Errorf("instruction format error: %s", instruction)
	}

	var action string
	var beginx, beginy, endX, endY int
	if parts[0] == "turn" {
		action = parts[1]
		start := strings.Split(parts[2], ",")
		end := strings.Split(parts[4], ",")

		if len(start) != 2 || len(end) != 2 {
			return fmt.Errorf("coordinate format error: %s", instruction)
		}

		var err error
		beginx, err = strconv.Atoi(start[0])
		if err != nil {
			fmt.Println("invalid beginx")
		}
		beginy, err = strconv.Atoi(start[1])
		if err != nil {
			fmt.Println("invalid beginy")
		}
		endX, err = strconv.Atoi(end[0])
		if err != nil {
			fmt.Println("invalid endx")
		}
		endY, err = strconv.Atoi(end[1])
		if err != nil {
			fmt.Println("invalid endy")
		}

	} else if parts[0] == "toggle" {
		action = "toggle"
		start := strings.Split(parts[1], ",")
		end := strings.Split(parts[3], ",")

		if len(start) != 2 || len(end) != 2 {
			return fmt.Errorf("coordinate format error: %s", instruction)
		}

		var err error
		beginx, err = strconv.Atoi(start[0])
		if err != nil {
			fmt.Println("invalid beginx")
		}
		beginy, err = strconv.Atoi(start[1])
		if err != nil {
			fmt.Println("invalid beginy")
		}
		endX, err = strconv.Atoi(end[0])
		if err != nil {
			fmt.Println("invalid endx")
		}
		endY, err = strconv.Atoi(end[1])
		if err != nil {
			fmt.Println("invalid endy")
		}
	} else {
		return fmt.Errorf("unknown action: %s", parts[0])
	}

	// Apply the instruction to the grid
	for x := beginx; x <= endX; x++ {
		for y := beginy; y <= endY; y++ {
			switch action {
			case "on":
				grid[x][y]++
			case "off":
				if grid[x][y] > 0 {
					grid[x][y]--
				}
			case "toggle":
				grid[x][y] += 2
			}
		}
	}
	return nil
}

// Calculates the total brightness of all lights
func calculateTotalBrightness(grid [][]int) int {
	totalBrightness := 0
	for i := range grid {
		for j := range grid[i] {
			totalBrightness += grid[i][j]
		}
	}
	return totalBrightness
}
