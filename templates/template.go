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


section .bss
    pointer: resb 680_000
    counter: resb 8


section .data
    true: db "true", 0xa
    trueLength equ $-true
    false: db "false", 0xa
    falseLength equ $-false
%STRINGS%

`, "%FUNCTIONS%", print+memory, 1)
