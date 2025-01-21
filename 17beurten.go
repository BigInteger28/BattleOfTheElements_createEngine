package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for {
		// Get input from the player
		var input string
		fmt.Print("Enter first engine depth(s): ")
		fmt.Scanln(&input)

		// Check the length of the input
		inputLength := len(input)
		if inputLength > 16 {
			fmt.Println("Input must be 16 digits or less.")
			continue
		}

		// If input length is less than 16, generate random numbers to fill
		extraNumbers := ""
		for i := 0; i < 16-inputLength; i++ {
			extraNumbers += fmt.Sprintf("%d", rand.Intn(4)+1)
		}

		// Concatenate input with extra random numbers
		result := input + extraNumbers

		// Output the result
		fmt.Println("Generated Output:", result)
		fmt.Println()
	}
}
