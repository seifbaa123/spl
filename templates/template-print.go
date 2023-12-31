package templates

var print = printInt + printChar + printBool + printStr

var printInt = `
_print_int:
    push rax
    push rdi
    push rsi
    push rdx
    push rcx
    push rbp

    mov rax, [rsp+56]
    mov rbp, rsp
    mov rcx, 2

    cmp rax, 0
    jl .negative_number

.positive_number:
    mov rdi, 0
    jmp .work

.negative_number:
    mov rdi, 1
    not rax
    add rax, 1        

.work:
    push byte 0x0
    push byte 0xa

.push:
    mov rdx, 0
    mov rbx, 10
    div rbx
    add rdx, 0x30
    push rdx
    add rcx, 8
    cmp rax, 0
    jne .push
    cmp rdi, 1
    jz .push_minus
    jmp .print

.push_minus:
    push '-'
    add rcx, 8

.print:
    mov rax, 1
    mov rdi, 1
    mov rsi, rsp
    mov rdx, rcx
    syscall

    mov rsp, rbp

    pop rbp
    pop rcx
    pop rdx
    pop rsi
    pop rdi
    pop rax

    ret
`

var printChar = `
_print_char:
    push rax
    push rdi
    push rsi
    push rdx

    mov rsi, [rsp+5*8]
    
    push 0xa
    push rsi

    mov rax, 1
    mov rdi, 1
    mov rsi, rsp
    mov rdx, 16
    syscall

    add rsp, 16

    pop rdx
    pop rsi
    pop rdi
    pop rax

    ret
`

var printBool = `
_print_bool:
    push rax
    push rdi
    push rsi
    push rdx

    mov rax, [rsp+5*8]

    cmp rax, 0
    je .print_false
    jmp .print_true

.print_true:
    mov rax, 1
    mov rdi, 1
    mov rsi, true
    mov rdx, trueLength
    syscall
    jmp .end

.print_false:
    mov rax, 1
    mov rdi, 1
    mov rsi, false
    mov rdx, falseLength
    syscall

.end:

    pop rdx
    pop rsi
    pop rdi
    pop rax

    ret
`

var printStr = `
_print_str:
    push rax
    push rdi
    push rsi
    push rdx
    push rbx

    mov rbx, [rsp+6*8]
    mov rbx, [rbx]

    mov rax, 1
    mov rdi, 1
    mov rdx, [rbx]
    add rbx, 8
    mov rsi, rbx
    syscall

    pop rbx
    pop rdx
    pop rsi
    pop rdi
    pop rax

    ret
`
