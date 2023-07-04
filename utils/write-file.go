package utils

import (
	"fmt"
	"os"
)

func WriteFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not write file %s!\n", filename)
		os.Exit(1)
	}
}
