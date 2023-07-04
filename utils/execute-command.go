package utils

import (
	"fmt"
	"os/exec"
)

func Execute(app string, args ...string) {
	cmd := exec.Command(app, args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(out))
}
