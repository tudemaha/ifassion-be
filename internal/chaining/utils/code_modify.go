package utils

import (
	"fmt"
	"strconv"
)

func IncrementCode(codeString string) string {
	num, _ := strconv.Atoi(codeString[1:])
	num++

	incrementedCodeString := fmt.Sprintf("%c%02d", codeString[0], num)
	return incrementedCodeString
}
