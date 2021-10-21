package set1

import (
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// var vowels = []string{
// 	"a",
// 	"e",
// 	"i",
// 	"o",
// 	"u",
// 	"A",
// 	"E",
// 	"I",
// 	"O",
// 	"U",
// }
var vowels = []string{
	"e",
	"t",
	"a",
	"o",
	"i",
	"n",
	"s",
	"E",
	"T",
	"A",
	"O",
	"I",
	"N",
	"S",
}

func SingleXORBrute(input string) error {
	hexStr, err := hex.DecodeString(input)
	if err != nil {
		return err
	}

	for i := 0; i < 255; i++ {
		fmt.Println(stringXOR(string(hexStr), string(i)))
	}

	return nil
}

func GuessXOR(input string) (string, error) {

	possibleAnswers := []string{}

	hexStr, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	highestmatchCount := 0
	currentAnswer := ""

	for i := 0; i < 255; i++ {
		result := (stringXOR(string(hexStr), string(i)))
		matchCount := 0

		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(result, "")

		// for x := 0; x < len(vowels); x++ {
		// 	matchCount += strings.Count(processedString, vowels[x])
		// }
		matchCount = GetWordScore(processedString)

		if matchCount >= highestmatchCount {
			highestmatchCount = matchCount
			currentAnswer = result
			possibleAnswers = append(possibleAnswers, result)
		}
	}

	return currentAnswer, nil
}

func GetWordScore(word string) int {
	matchCount := 0
	for x := 0; x < len(vowels); x++ {
		matchCount += strings.Count(word, vowels[x])
	}

	return matchCount
}
