package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for {
		var amount, startseed int
		var input, extraNumbers, result string
		fmt.Print("Start Number: ")
		fmt.Scanln(&startseed)
		fmt.Print("Amount of generations: ")
		fmt.Scanln(&amount)
		fmt.Print("Enter first engine depth(s): ")
		fmt.Scanln(&input)
		
		// Check the length of the input
		inputLength := len(input)
		if inputLength > 12 {
			fmt.Println("Input must be 12 digits or less.")
			continue
		}
		
		for i := 0; i < amount; i++ {
			// If input length is less than 12, generate random numbers to fill
			extraNumbers = ""
			for j := 0; j < 12-inputLength; j++ {
				extraNumbers += fmt.Sprintf("%d", rand.Intn(4)+1)
			}
			// Concatenate input with extra random numbers
			result = "x" + strconv.Itoa(startseed) + ":1:" + input + extraNumbers
			fmt.Println(result)
			startseed++
		}		
		fmt.Println()
	}
}
