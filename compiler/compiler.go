package compiler

import (
	"fmt"
	"spl/statements"
	"spl/templates"
	"spl/utils"
	"strings"
)

func Compile(program *statements.Program) {
	code := strings.Replace(templates.Main, "%CODE%", program.Evaluate().Assembly, 1)

	r := utils.RandomString(20)
	assemblyFile := fmt.Sprintf("/tmp/code.%s.asm", r)
	utils.WriteFile(assemblyFile, code)

	objectFile := fmt.Sprintf("/tmp/res.%s.o", r)
	utils.Execute("nasm", "-f", "elf64", "-o", objectFile, assemblyFile)

	executableFile := fmt.Sprintf("/tmp/res.%s", r)
	utils.Execute("ld", "-o", executableFile, objectFile)

	utils.Execute(executableFile)
}
