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

func Or(val1 string, val2 string) string {
	return fmt.Sprintf("    or %s, %s", val1, val2)
}

func And(val1 string, val2 string) string {
	return fmt.Sprintf("    and %s, %s", val1, val2)
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

func Cmp(val1 string, val2 string) string {
	return fmt.Sprintf("    cmp %s, %s", val1, val2)
}

func Jmp(val string) string {
	return fmt.Sprintf("    jmp %s", val)
}

func Je(val string) string {
	return fmt.Sprintf("    je %s", val)
}

func Jne(val string) string {
	return fmt.Sprintf("    jne %s", val)
}

func Jg(val string) string {
	return fmt.Sprintf("    jg %s", val)
}

func Jge(val string) string {
	return fmt.Sprintf("    jge %s", val)
}

func Jl(val string) string {
	return fmt.Sprintf("    jl %s", val)
}

func Jle(val string) string {
	return fmt.Sprintf("    jle %s", val)
}
