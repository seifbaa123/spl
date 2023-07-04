package instructions

import "fmt"

func Mov(val1 string, val2 string) string {
	return fmt.Sprintf("    mov %s, %s", val1, val2)
}

func Add(val1 string, val2 string) string {
	return fmt.Sprintf("    add %s, %s", val1, val2)
}

func Sub(val1 string, val2 string) string {
	return fmt.Sprintf("    sub %s, %s", val1, val2)
}

func Mul(val string) string {
	return fmt.Sprintf("    mul %s", val)
}

func Div(val string) string {
	return fmt.Sprintf("    div %s", val)
}

func Xor(val1 string, val2 string) string {
	return fmt.Sprintf("    xor %s, %s", val1, val2)
}

func Push(val string) string {
	return fmt.Sprintf("    push %s", val)
}

func Pop(val string) string {
	return fmt.Sprintf("    pop %s", val)
}

func Call(val string) string {
	return fmt.Sprintf("    call %s", val)
}
