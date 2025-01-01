//guessing_game.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(100) + 1

	fmt.Println("Guess a number between 1 and 100")

	var guess int

	for {
		fmt.Print("Enter your guess:")
		fmt.Scanln(&guess)

		//check
		if guess < target {
			fmt.Println("Too low")
		} else if guess > target {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it!")
			break
		}
	}
}