package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	turns = 5
	usage = `Welcome to lucky numbers! Wanna play? The program will pick %d random numbers, Pick two positive numbers`
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf(usage, turns)
		return
	}
	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Printf("Not two numbers")
		return
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Printf("Please pick positive numbers")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for turn := 0; turn <= turns; turn++ {

		n := rand.Intn(guess1 + 1)
		if n == guess1 || n == guess2 {

			if turn == 0 {
				fmt.Printf("Congratulations, you won in first go!")
				return
			}
			fmt.Printf("You win in %d turn!", turn+1)
			return
		}
	}
	fmt.Printf("You lost! Try again")
}
