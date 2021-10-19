package set1

import (
	"testing"
)

var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

var output string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func TestConvert(t *testing.T) {
	result, _ := Convert(input)
	if result != output {
		t.Fatalf("Result: %s Expected: %s", result, output)
	}
}
