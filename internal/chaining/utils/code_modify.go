package utils

import (
	"fmt"
	"sort"
	"strconv"
)

func IncrementCode(sliceTrue, sliceFalse []string) string {
	sliceNum := make([]int, 0)

	for _, element := range sliceTrue {
		num, _ := strconv.Atoi(element[1:])
		sliceNum = append(sliceNum, num)
	}

	for _, element := range sliceFalse {
		num, _ := strconv.Atoi(element[1:])
		sliceNum = append(sliceNum, num)
	}

	smallest := 1
	sort.Ints(sliceNum)
	for _, element := range sliceNum {
		fmt.Println(element)
		if element == smallest {
			smallest++
		} else {
			break
		}
	}

	incrementedCodeString := fmt.Sprintf("%c%02d", 'I', smallest)
	return incrementedCodeString
}
