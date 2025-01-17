.section .text
.globl _start

.equ UART_BASE,  0x10013000
.equ GPIO_BASE,  0x10012000
.equ GPIO_OUTPUT_EN,  0x08
.equ GPIO_OUTPUT_VAL, 0x0c

#.equ PIN7, 1 # pin no and gpio no are different
.equ PIN1, 1

_start:

# s0 holds GPIO structure address
la s0, GPIO_BASE

# create pin bitmask
# stored in s1
li t0, PIN1
slli s1, t0, 1

# set pin as output
sw s1, GPIO_OUTPUT_EN(s0)
