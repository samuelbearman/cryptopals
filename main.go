package main

import (
	"fmt"

	"github.com/samuelbearman/cryptopals/set1"
)

func main() {
	runSet1()
}

func runSet1() {
	var input string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	answer, _ := set1.GuessXOR(input)
	fmt.Printf("Best Guess: %s", answer)
}
