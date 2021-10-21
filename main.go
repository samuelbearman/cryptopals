package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/samuelbearman/cryptopals/set1"
)

func main() {
	// runSet1()
	runSet1Challenge4()
}

func runSet1() {
	// var input string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	// answer, _ := set1.GuessXOR(input)
	// fmt.Printf("Best Guess: %s", answer)
}

func runSet1Challenge4() {
	potentialAnswers := []string{}

	file, err := os.Open("data/challenge4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result, _ := set1.GuessXOR(scanner.Text())
		// fmt.Println(result)
		potentialAnswers = append(potentialAnswers, result)
	}

	matchCount := 0
	bestGuess := ""

	for _, v := range potentialAnswers {
		score := set1.GetWordScore(v)

		if score >= matchCount {
			bestGuess = v
			matchCount = score
		}
		// fmt.Println(v)
	}

	fmt.Println(bestGuess)

}
