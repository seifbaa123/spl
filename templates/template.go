package templates

import "strings"

var Main = strings.Replace(`BITS 64
global _start

section .text
_start:
    mov rbp, rsp

%CODE%

    mov rax, 60
    mov rbx, 0
    syscall

%FUNCTIONS%

section .data
    true: db "true", 0xa
    trueLength equ $-true
    false: db "false", 0xa
    falseLength equ $-false

`, "%FUNCTIONS%", printInt+printChar+printBool, 1)
