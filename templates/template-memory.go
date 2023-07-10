package templates

var memory = allocate + free + memCopy

var allocate = `
_allocate:
    push rax
    push rbx
    push rcx

    mov rbx, [rsp+8*4]

    mov rax, [counter]
    mov rcx, rax
    add rcx, rbx
    mov [counter], rcx
    add rax, pointer

    pop rcx
    pop rbx
    add rsp, 8

    ret
`

var free = `
_free:
    ret
`

var memCopy = `
_mem_copy:
    push rax
    push rbx
    push rcx
    push rdx

    mov rax, [rsp+8*7] ; from address
    mov rbx, [rsp+8*6] ; to address
    mov rcx, [rsp+8*5] ; size

.loop:
    cmp rcx, 0
    jz .end
        
    dec rcx
    mov dl, [rax+rcx]
    mov [rbx+rcx], dl
        
    jmp .loop

.end:

    pop rdx
    pop rcx
    pop rbx
    pop rax

    ret
`
