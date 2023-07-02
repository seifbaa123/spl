package utils

import (
	"fmt"
	"os"
)

func ReadSource() string {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [file.spl]\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	src, err := os.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file %s\n", filename)
		os.Exit(1)
	}

	return string(src)
}
