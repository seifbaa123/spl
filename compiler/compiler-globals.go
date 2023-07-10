package compiler

import (
	"fmt"
	"strings"
)

var stringsList []string

func addString(str string) int {
	num := len(stringsList)
	stringsList = append(stringsList, str)

	return num
}

func getStrings() string {
	var str []string

	for i, s := range stringsList {
		str = append(str, fmt.Sprintf("    str%d: db \"%s\"", i, s))
	}

	return strings.Join(str, "\n")
}
