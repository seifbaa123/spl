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
		str = append(str, fmt.Sprintf("    str%d: db %s", i, joinString(s)))
	}

	return strings.Join(str, "\n")
}

func joinString(str string) string {
	var s []string

	for _, c := range str {
		s = append(s, fmt.Sprint(int(c)))
	}

	return strings.Join(s, ", ")
}
