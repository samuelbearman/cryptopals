package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strings"
)

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

func RepeatingKeyXOR(input, key []byte) []byte {
	eb := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		eb[i] = input[i] ^ key[i%len(key)]
		fmt.Printf("%x, %x, %x\n", input[i], key[i%len(key)], eb[i])
	}
	return eb
}

func Encode(plainText []byte) []byte {
	ret := make([]byte, hex.EncodedLen(len(plainText)))
	hex.Encode(ret, plainText)
	return ret
}

func HexBase64Convert(hexstr string) (string, error) {
	asciiStr, err := hex.DecodeString(hexstr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString((asciiStr)), nil
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

func FixedXORConvert(buf1 string, buf2 string) (string, error) {
	hexStr, err := hex.DecodeString(buf1)
	if err != nil {
		return "", err
	}

	result := hex.EncodeToString([]byte(stringXOR(string(hexStr), buf2)))
	return result, nil
}

func stringXOR(value string, key string) (output string) {
	for i := 0; i < len(value); i++ {
		output += string(value[i] ^ key[i%len(key)])
	}

	return output
}
