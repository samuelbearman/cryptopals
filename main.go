package main

import (
	"fmt"

	"github.com/samuelbearman/cryptopals/set1"
)

func main() {
	runSet1()
}

func runSet1() {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	result, err := set1.Convert(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
