package set1Test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/samuelbearman/cryptopals/set1"
)

func TestHexBase64(t *testing.T) {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, _ := set1.Challenge1_HexBase64Convert(input)
	if result != output {
		t.Fatalf("Result: %s Expected: %s", result, output)
	}
}

func TestFixedXOR(t *testing.T) {
	output := "746865206b696420646f6e277420706c6179"
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	result, _ := set1.Challenge2_FixedXORConvert(input1, input2)
	if result != output {
		t.Fatalf("Result: %s Expected: %s", result, output)
	}
}

func TestSingleByteXOR(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// result := set1.SingleXORBrute(input)
	// if result != nil {
	// 	t.Fatalf("Error: %s", result)
	// }

	result, _ := set1.Challenge3_GuessXOR(input)
	fmt.Printf("Result: %s", result)
}

func TestSingleByteXORFromFile(t *testing.T) {
	result := set1.Challenge4_DetectXORFromFile("../data/challenge4.txt")
	fmt.Printf("Result: %s", result)
}

func TestRepeatingKeyXOR(t *testing.T) {
	input := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	result := set1.Challenge5_RepeatingKeyXOR(input, []byte("ICE"))

	expected := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	encodedResult := set1.Encode(result)

	if !bytes.Equal(encodedResult, expected) {
		t.Fatal("Not Equal")
	}

}

func TestBreakRepeatingKeyXOR(t *testing.T) {
	editDistance, _ := set1.CalculateEditDistance("this is a test", "wokka wokka!!!")

	if editDistance != 37 {
		t.Errorf("Expected: %d Result: %d", 37, editDistance)
	}

}
