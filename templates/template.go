package templates

var Main = `BITS 64
global _start

section .text
_start:
%CODE%

    mov rax, 60
    mov rbx, 0
    syscall
` + printInt + printChar
