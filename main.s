.section .text
.globl _start

.equ UART_BASE,  0x10013000
.equ GPIO_BASE,  0x10012000
.equ GPIO_OUTPUT_EN,  0x08
.equ GPIO_OUTPUT_VAL, 0x0c

.equ PIN1, 1

_start:
la s0, GPIO_BASE

# create pin bitmask
# stored in s1
li t0, PIN1
slli s1, t0, 1

# set pin as output
sw s1, GPIO_OUTPUT_EN(s0)

# inst prerequisite
li t0, 1
li t1, 2

# loop n times
li s2, 100
loop:

# set pin high
sw s1, GPIO_OUTPUT_VAL(s0)

add t3, t0, t1

# set pin low
li t0, 0x0
sw t0, GPIO_OUTPUT_VAL(s0)

# print
la a0, h
jal print_char


# decrease loop counter
addi s2, s2, -1
bnez s2, loop


halt: j halt



print_char:
	li t5, UART_BASE
	lb a1, (a0)
	beqz a1, end_print
wait:
	lw t6, (t5)
	bltz t6, wait

	sw a1, (t5)
	addi a0, a0, 1
	j print_char

end_print:
	ret

.section .rodata
h: .string "h"
