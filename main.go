package main

import (
	"fmt"
)

func divisibleByThree(number int) bool {
	return ((number % 3) == 0)
}

func divisibleByFive(number int) bool {
	return ((number % 5) == 0)
}

func main() {
	fmt.Println("Challenge 1 =====================")

	for i := 1; i <= 100; i++ {
		if divisibleByThree(i) {
			if i != 100 {
				fmt.Printf("%d, ", i)
			} else {
				fmt.Printf("%d", i)
			}

		}
	}

	fmt.Println("\n\nChallenge 2 =====================")

	for i := 1; i <= 100; i++ {
		if divisibleByThree(i) && divisibleByFive(i) {
			fmt.Println(i, " - PinPan")
		} else if divisibleByThree(i) {
			fmt.Println(i, " - Pin")
		} else if divisibleByFive(i) {
			fmt.Println(i, " - Pan")
		}
	}
}
