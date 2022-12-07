package subsequence

import (
	"reflect"
	"strings"
)

func ValidSubsequence(seq, str string) bool {
	resArr := make([]string, 0)

	splitStr := strings.Split(str, "")
	splitSeq := strings.Split(seq, "")

	if len(splitSeq) > len(splitStr) {
		return false
	}

	pointer := 0

	for i, _ := range splitStr {
		if splitStr[i] == splitSeq[pointer] {
			resArr = append(resArr, splitStr[i])
			pointer++

			if reflect.DeepEqual(resArr, splitSeq) {
				return true
			}
		}
	}

	return false
}
