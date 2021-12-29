package set1

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var mostCommonLetters = []string{
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

func Challenge1_HexBase64Convert(hexstr string) (string, error) {
	asciiStr, err := hex.DecodeString(hexstr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString((asciiStr)), nil
}

func Challenge2_FixedXORConvert(buf1 string, buf2 string) (string, error) {
	hexStr, err := hex.DecodeString(buf1)
	if err != nil {
		return "", err
	}

	result := hex.EncodeToString([]byte(stringXOR(string(hexStr), buf2)))
	return result, nil
}

func Challenge3_GuessXOR(input string) (string, error) {

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
		matchCount = getWordScore(processedString)

		if matchCount >= highestmatchCount {
			highestmatchCount = matchCount
			currentAnswer = result
			possibleAnswers = append(possibleAnswers, result)
		}
	}

	return currentAnswer, nil
}

func Challenge4_DetectXORFromFile(path string) string {
	potentialAnswers := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result, _ := Challenge3_GuessXOR(scanner.Text())
		potentialAnswers = append(potentialAnswers, result)
	}

	matchCount := 0
	bestGuess := ""

	for _, v := range potentialAnswers {
		score := getWordScore(v)

		if score >= matchCount {
			bestGuess = v
			matchCount = score
		}
	}
	return bestGuess
}

func SingleXORBrute(input string) error {
	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		return err
	}

	hexStr := string(hexBytes)

	for i := 65; i < 123; i++ {
		letter := string(i)
		output := stringXOR(hexStr, letter)
		fmt.Printf("Letter: %s\n Output: %s \n\n ", letter, output)
	}

	return nil
}

func Challenge5_RepeatingKeyXOR(input, key []byte) []byte {
	eb := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		eb[i] = input[i] ^ key[i%len(key)]
		// fmt.Printf("%x, %x, %x\n", input[i], key[i%len(key)], eb[i])
	}
	return eb
}

func Challenge6_BreakRepeatingKeyXOR() {

}

func CalculateEditDistance(str1 string, str2 string) (int, error) {
	if len(str1) != len(str2) {
		return 0, errors.New("Length of strings do not match")
	}

	runningEditDistance := 0

	bytes1 := []byte(str1)
	bytes2 := []byte(str2)

	for x := 0; x < len(bytes1); x++ {
		firstByte := bytes1[x]
		secondByte := bytes2[x]

		diff := diff(firstByte, secondByte)
		runningEditDistance += diff
		fmt.Printf("Distance: %d\n", diff)
	}

	return runningEditDistance, nil
}

func diff(a, b byte) int {
	if a < b {
		return int(b - a)
	}
	return int(a - b)
}

func Encode(plainText []byte) []byte {
	ret := make([]byte, hex.EncodedLen(len(plainText)))
	hex.Encode(ret, plainText)
	return ret
}

func getWordScore(word string) int {
	matchCount := 0
	for x := 0; x < len(mostCommonLetters); x++ {
		matchCount += strings.Count(word, mostCommonLetters[x])
	}

	return matchCount
}

func stringXOR(value string, key string) (output string) {
	for i := 0; i < len(value); i++ {
		output += string(value[i] ^ key[i%len(key)])
	}

	return output
}
