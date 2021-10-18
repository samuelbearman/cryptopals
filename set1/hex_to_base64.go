package set1

import (
	"fmt"
)

func Convert(hexstr string) (string, error) {
	strByteArr := []byte(hexstr)

	fmt.Println(strByteArr)
	return "", nil
}
