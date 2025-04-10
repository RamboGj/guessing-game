package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing game")
	fmt.Println("A random number will be drawn. It's an integer between 0 and 100")

	x := rand.Int64N(101)
	x = 10
	// Create a scanner that reads from Stdin(OS reader)
	scanner := bufio.NewScanner(os.Stdin)
	guesses:= [10]int64{}

	for i:= range guesses {
		fmt.Printf("What is your guess? ")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, error := strconv.ParseInt(guess, 10, 64)

		if error != nil {
			fmt.Printf("Your guess must be an intenger number")
			return 
		}

		switch {
		case guessInt < x:
			fmt.Println("You did it wrong. The number is higher than", guessInt)
		case guessInt > x:
			fmt.Println("You did it wrong. The number is lower than", guessInt)
		case guessInt == x:
			fmt.Printf(
				"Congratulations, the number is: %d\n"+
				"You did it in %d attempts.\n"+
				"Take a look at your attemps: %v",
    			x, i+1, guesses[:i],
			)
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf(
		"Unfortunately, you could not guess the number. It was: %d\n" + 
		"You had 10 attempts: %v\n",
		x, guesses,
	)
}