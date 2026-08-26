package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Single global scanner to avoid buffer synchronization issues
var globalScanner = bufio.NewScanner(os.Stdin)

// Helper function to read a clean line from standard input
func readLine() (string, error) {
	if globalScanner.Scan() {
		return strings.TrimSpace(globalScanner.Text()), nil
	}
	if err := globalScanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("EOF")
}

// Scanln keeps looping until a line is successfully read (can be empty)
func Scanln() string {
	for {
		input, err := readLine()
		if err != nil {
			fmt.Print("Reading error. Please try again: ")
			continue
		}
		return input
	}
}

// Nextint keeps looping until the user enters a valid integer
func Nextint() int {
	for {
		input, err := readLine()
		if err != nil {
			fmt.Print("Reading error. Please try again: ")
			continue
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Error: not a valid integer. Please try again: ")
			continue
		}

		return value
	}
}

// Nextstring keeps looping until the user enters a non-empty string
func Nextstring() string {
	for {
		input, err := readLine()
		if err != nil {
			fmt.Print("Reading error. Please try again: ")
			continue
		}
		
		if input == "" {
			fmt.Print("Error: string cannot be empty. Please try again: ")
			continue
		}

		return input
	}
}

// Nextfloat32 keeps looping until the user enters a valid float32
func Nextfloat32() float32 {
	for {
		input, err := readLine()
		if err != nil {
			fmt.Print("Reading error. Please try again: ")
			continue
		}

		value, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Print("Error: not a valid decimal number. Please try again: ")
			continue
		}

		return float32(value)
	}
}

// Nextbool keeps looping until a boolean is entered (accepts true/false, 1/0, y/n, yes/no)
func Nextbool() bool {
	for {
		input, err := readLine()
		if err != nil {
			fmt.Print("Reading error. Please try again: ")
			continue
		}

		inputLower := strings.ToLower(input)

		// Map custom yes/no shortcuts to standard Go boolean formats
		if inputLower == "y" || inputLower == "yes" {
			input = "true"
		} else if inputLower == "n" || inputLower == "no" {
			input = "false"
		}

		value, err := strconv.ParseBool(input)
		if err != nil {
			fmt.Print("Error: invalid input (use true/false, 1/0, or y/n). Please try again: ")
			continue
		}

		return value
	}
}

