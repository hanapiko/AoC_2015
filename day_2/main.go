package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Open the data.txt file
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening file")
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    WrappingPaper := 0
    ribbon := 0

    for scanner.Scan() {
        line := scanner.Text()
        dimensions := strings.Split(line, "x")
        if len(dimensions) != 3 {
            fmt.Println("Invalid input format")
            continue
        }
        l, err := strconv.Atoi(dimensions[0])
        if err != nil {
            fmt.Println("Error parsing")
            continue
        }
        w, err := strconv.Atoi(dimensions[1])
        if err != nil {
            fmt.Println("Error parsing width:", err)
            continue
        }
        h, err := strconv.Atoi(dimensions[2])
        if err != nil {
            fmt.Println("Error parsing height:", err)
            continue
        }

        area1 := l * w
        area2 := w * h
        area3 := h * l

        sArea := 2*(area1 + area2 + area3)
        small:= min(area1, min(area2, area3))

        WrappingPaper += sArea + small

        perimeter1 := 2 * (l + w)
        perimeter2 := 2 * (w + h)
        perimeter3 := 2 * (h + l)
        smallestPerimeter := min(perimeter1, min(perimeter2, perimeter3))

        volume := l * w * h
        ribbon += smallestPerimeter + volume
    }

    fmt.Println(WrappingPaper)
    fmt.Println(ribbon)
}

// minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

