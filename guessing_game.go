// guessing_game.go
package guessinggame

import (
    "fmt"
    "math/rand"
    "time"
)

// Play starts the guessing game
func Play() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1

    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I have selected a random number between 1 and 100.")
    fmt.Println("Can you guess it?")

    var guess int
    for {
        fmt.Print("Enter your guess: ")
        fmt.Scanln(&guess)

        if guess < target {
            fmt.Println("Too low!")
        } else if guess > target {
            fmt.Println("Too high!")
        } else {
            fmt.Println("Congratulations! You guessed it!")
            break
        }
    }
}
