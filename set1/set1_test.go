package set_one

import (
	"bytes"
	"testing"
)

func Test_Challenge1_HexBase64(t *testing.T) {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, _ := Challenge1_HexBase64Convert(input)
	if result != expected {
		t.Fatalf("Result: %s Expected: %s", result, expected)
	}
}

func Test_Challenge2_FixedXOR(t *testing.T) {
	expected := "746865206b696420646f6e277420706c6179"
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	result, _ := Challenge2_FixedXORConvert(input1, input2)
	if result != expected {
		t.Fatalf("Result: %s Expected: %s", result, expected)
	}
}

func Test_Challenge3_SingleByteXOR(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	result, _ := Challenge3_GuessXOR(input)
	t.Logf("Challenge 3: %s", result)
}

func Test_Challenge4_SingleByteXORFromFile(t *testing.T) {
	result := Challenge4_DetectXORFromFile("../data/challenge4.txt")
	t.Logf("Challenge 4: %s", result)
}

func Test_Challenge5_RepeatingKeyXOR(t *testing.T) {
	input := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	result := Challenge5_RepeatingKeyXOR(input, []byte("ICE"))

	expected := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	encodedResult := Encode(result)

	if !bytes.Equal(encodedResult, expected) {
		t.Fatalf("Result: %s Expected: %s", encodedResult, expected)
	}

}
