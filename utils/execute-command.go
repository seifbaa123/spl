package utils

import (
	"os"
	"os/exec"
)

func Execute(app string, args ...string) {
	cmd := exec.Command(app, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
