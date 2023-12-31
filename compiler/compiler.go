package compiler

import (
	"fmt"
	"spl/templates"
	"spl/utils"
	"strings"
)

func Compile(program *Program) {
	env := NewEnvironment()
	code := templates.Main

	code = strings.Replace(code, "%CODE%", program.Evaluate(env).Assembly, 1)
	code = strings.Replace(code, "%STRINGS%", getStrings(), 1)

	r := utils.RandomString(20)
	assemblyFile := fmt.Sprintf("/tmp/code.%s.asm", r)
	utils.WriteFile(assemblyFile, code)

	objectFile := fmt.Sprintf("/tmp/res.%s.o", r)
	utils.Execute("nasm", "-f", "elf64", "-o", objectFile, assemblyFile)

	executableFile := fmt.Sprintf("/tmp/res.%s", r)
	utils.Execute("ld", "-o", executableFile, objectFile)

	utils.Execute(executableFile)
}
