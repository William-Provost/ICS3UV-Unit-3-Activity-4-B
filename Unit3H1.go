/*
 * Author: William Provost
 * Version: 1.0.0
 * Date: 2025-11-13
 * Fileoverview: This program tells you what to do in different weather.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var chanceOfRainAsString string = ""
	var chanceOfRainAsNumber int = 0

	var reader = bufio.NewReader(os.Stdin)

	// input % chance of rain
	fmt.Print("Enter the chance of rain (0-100): ")
	chanceOfRainAsString, _ = reader.ReadString('\n')
	chanceOfRainAsString = strings.TrimSpace(chanceOfRainAsString)
	chanceOfRainAsNumber, _ = strconv.Atoi(chanceOfRainAsString)

	// check the chance of rain
	if chanceOfRainAsNumber > 80 {
		// what to bring if raining
		fmt.Println("Bring an umbrella.")
		fmt.Println("Also bring rain boots!")
	} else {
		// what to bring if it is sunny
		fmt.Println("Bring a sun hat.")
		fmt.Println("Also bring a swimsuit!")
	}

	fmt.Println("\nDone.")
}
