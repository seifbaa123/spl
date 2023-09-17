package templates

var str = strCompare + checkIndex

var strCompare = `
_str_compare:
    push rax
    push rsi
    push rdi
    push rdx
    push rcx

    mov rax, [rsp+8*6]
    mov rdx, [rsp+8*7]

    mov rax, [rax]
    mov rdx, [rdx]

    ; compare length
    mov rsi, [rax]
    mov rdi, [rdx]

    cmp rsi, rdi
    jne .false

    ; compare the characters
    mov rcx, rsi
    
.compare:
    cmp rcx, 0
    jz .true

    dec rcx
    movzx rsi, byte [rax+rcx+8]
    movzx rdi, byte [rdx+rcx+8]
        
    cmp rsi, rdi
    jne .false
    jmp .compare

.false:
    xor rax, rax
    jmp .end
    
.true:
    mov rax, 1

.end:
    
    pop rcx
    pop rdx
    pop rdi
    pop rsi
    add rsp, 8

    ret
`

var checkIndex = `
_check_index:
    push rax    
    push rdi
    
    mov rax, [rsp+4*8]
    mov rdi, [rsp+3*8]

    cmp rdi, 0
    jl .print_error

    cmp rdi, rax
    jge .print_error

    pop rdi
    pop rax

    ret

.print_error:
    mov rax, 1
    mov rdi, 1
    mov rsi, indexOutOfRange
    mov rdx, indexOutOfRangeLength
    syscall

    mov rax, 60
    mov rdi, 1
    syscall
`
