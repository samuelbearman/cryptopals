package set1

import (
	"testing"
)

func TestHexBase64(t *testing.T) {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, _ := HexBase64Convert(input)
	if result != output {
		t.Fatalf("Result: %s Expected: %s", result, output)
	}
}

func TestFixedXOR(t *testing.T) {
	output := "746865206b696420646f6e277420706c6179"
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	result, _ := FixedXORConvert(input1, input2)
	if result != output {
		t.Fatalf("Result: %s Expected: %s", result, output)
	}
}
