package main

import (
	"os"
	"fmt"
	"AoC_2015/day_1/floor"
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	file, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("error opening file")
	}

	dest :=  floor.FloorNum(string(file))
	fmt.Println(dest)
}