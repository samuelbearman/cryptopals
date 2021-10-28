package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/samuelbearman/cryptopals/set1"
)

func main() {
	// runSet1()
	// runSet1Challenge4()
	runSet1Challenge5()
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
func runSet1Challenge5() {
	input := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	result := set1.RepeatingKeyXOR(input, []byte("ICE"))
	fmt.Println(result)
	// fmt.Println(input)
	expected := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	// fmt.Println(expected)
	encodedResult := set1.Encode(result)

	if !bytes.Equal(encodedResult, expected) {
		fmt.Println("Not equal")
	}

}

func Encode(plaintext []byte) []byte {
	ret := make([]byte, base64.StdEncoding.EncodedLen(len(plaintext)))
	base64.StdEncoding.Encode(ret, plaintext)
	return ret
}
