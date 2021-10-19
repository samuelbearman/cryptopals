package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func Convert(hexstr string) (string, error) {
	asciiStr, err := hex.DecodeString(hexstr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString((asciiStr)), nil
}
