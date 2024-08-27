package main

import (
	"fmt"
	"os"

	"AoC_2015/day_1/floor"
)

func main() {
	// welcome
	if len(os.Args) != 1 {
		return
	}
	file, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("error opening file")
	}

	dest := floor.FloorNum(string(file))
	Position := floor.Position(string(file))

	fmt.Println(dest)
	fmt.Println(Position)
}
