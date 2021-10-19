package set1

import (
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var vowels = []string{
	"a",
	"e",
	"i",
	"o",
	"u",
	"A",
	"E",
	"I",
	"O",
	"U",
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

	highestVowelCount := 0
	currentAnswer := ""

	for i := 0; i < 255; i++ {
		result := (stringXOR(string(hexStr), string(i)))
		vowelCount := 0

		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(result, "")

		for x := 0; x < len(vowels); x++ {
			vowelCount += strings.Count(processedString, vowels[x])
		}

		if vowelCount >= highestVowelCount {
			highestVowelCount = vowelCount
			currentAnswer = result
			possibleAnswers = append(possibleAnswers, result)
		}
	}

	fmt.Println("Possible Answers")

	for _, element := range possibleAnswers {
		fmt.Println(element)
	}

	return currentAnswer, nil
}
