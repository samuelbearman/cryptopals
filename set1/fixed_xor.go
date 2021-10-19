package set1

import "encoding/hex"

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
