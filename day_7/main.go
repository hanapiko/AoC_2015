package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//AND operation
func bitwiseAnd(val1, val2 uint16) uint16 {
	return val1 & val2
}

//OR operation
func bitwiseOr(val1, val2 uint16) uint16 {
	return val1 | val2
}

// Left shift operation
func leftShift(val1, val2 uint16) uint16 {
	return val1 << val2
}

// Right shift operation
func rightShift(val1, val2 uint16) uint16 {
	return val1 >> val2
}

// NOT operation
func bitwiseNot(val2 uint16) uint16 {
	return ^val2 & 0xFFFF // Mask to 16 bits
}

func assign(val1 uint16) uint16 {
	return val1
}

// Instruction represents an operation to be executed
type Instruction struct {
	Operation string
	Val1      string
	Val2      string
	Wire      string
}

// ExecuteInstruction performs the operation based on the instruction
func executeInstruction(inst Instruction, values map[string]uint16) {
	var val1, val2 uint16
	var ok bool

	// Handle val1
	if inst.Val1 != "" {
		if val1, ok = values[inst.Val1]; !ok {
			if numVal, err := strconv.ParseUint(inst.Val1, 10, 16); err == nil {
				val1 = uint16(numVal)
			} else {
				return
			}
		}
	}

	// Handle val2
	if inst.Val2 != "" {
		if val2, ok = values[inst.Val2]; !ok {
			if numVal, err := strconv.ParseUint(inst.Val2, 10, 16); err == nil {
				val2 = uint16(numVal)
			} else {
				return
			}
		}
	}

	switch inst.Operation {
	case "AND":
		values[inst.Wire] = bitwiseAnd(val1, val2)
	case "OR":
		values[inst.Wire] = bitwiseOr(val1, val2)
	case "LSHIFT":
		values[inst.Wire] = leftShift(val1, val2)
	case "RSHIFT":
		values[inst.Wire] = rightShift(val1, val2)
	case "NOT":
		values[inst.Wire] = bitwiseNot(val1) // NOT only needs one operand (val1)
	case "ASSIGN":
		values[inst.Wire] = assign(val1)
	}
}

// ParseInput reads the instructions from the file
func parseInput(filename string) ([]Instruction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []Instruction

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		input := parts[0]
		wire := parts[1]

		var op1, op2, op string

		if strings.Contains(input, " AND ") {
			parts2 := strings.Split(input, " AND ")
			if len(parts2) != 2 {
				return nil, fmt.Errorf("invalid AND operation format: %s", line)
			}
			op1 = parts2[0]
			op2 = parts2[1]
			op = "AND"
		} else if strings.Contains(input, " OR ") {
			parts2 := strings.Split(input, " OR ")
			if len(parts2) != 2 {
				return nil, fmt.Errorf("invalid OR operation format: %s", line)
			}
			op1 = parts2[0]
			op2 = parts2[1]
			op = "OR"
		} else if strings.Contains(input, " LSHIFT ") {
			parts2 := strings.Split(input, " LSHIFT ")
			if len(parts2) != 2 {
				return nil, fmt.Errorf("invalid LSHIFT operation format: %s", line)
			}
			op1 = parts2[0]
			op2 = parts2[1]
			op = "LSHIFT"
		} else if strings.Contains(input, " RSHIFT ") {
			parts2 := strings.Split(input, " RSHIFT ")
			if len(parts2) != 2 {
				return nil, fmt.Errorf("invalid RSHIFT operation format: %s", line)
			}
			op1 = parts2[0]
			op2 = parts2[1]
			op = "RSHIFT"
		} else if strings.HasPrefix(input, "NOT ") {
			parts2 := strings.TrimPrefix(input, "NOT ")
			op1 = parts2
			op = "NOT"
		} else {
			// Assuming an assignment operation if no known operator is found
			op1 = input
			op = "ASSIGN"
		}

		instructions = append(instructions, Instruction{
			Operation: op,
			Val1:      op1,
			Val2:      op2,
			Wire:      wire,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}

// EvaluateCircuit processes instructions and computes the results
func evaluateCircuit(instructions []Instruction) map[string]uint16 {
	values := make(map[string]uint16)

	for {
		updated := false
		for _, inst := range instructions {
			if _, ok := values[inst.Wire]; !ok {
				executeInstruction(inst, values)
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	return values
}

func partOne(filename string) (uint16, error) {
	instructions, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	values := evaluateCircuit(instructions)
	return values["a"], nil
}

func partTwo(filename string) (uint16, error) {
	instructions, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	// Get the value of wire 'a' from part one
	aSignal, err := partOne(filename)
	if err != nil {
		return 0, err
	}

	// Override wire 'b' with the signal from wire 'a'
	for i, inst := range instructions {
		if inst.Wire == "b" {
			instructions[i].Val1 = strconv.Itoa(int(aSignal))
			instructions[i].Val2 = ""
			instructions[i].Operation = "ASSIGN"
		}
	}

	// Reset all other wires
	values := make(map[string]uint16)

	// Re-run the circuit evaluation
	values = evaluateCircuit(instructions)

	return values["a"], nil
}

func main() {
	// Part one
	aSignal, err := partOne("data.txt")
	if err != nil {
		fmt.Println("Error in part one:", err)
		return
	}
	fmt.Printf("Signal for wire a in part one: %d\n", aSignal)

	// Part two
	newASignal, err := partTwo("data.txt")
	if err != nil {
		fmt.Println("Error in part two:", err)
		return
	}
	fmt.Printf("Signal for wire a in part two: %d\n", newASignal)
}
