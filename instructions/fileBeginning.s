.section .text
.globl _start

.equ UART_BASE,  0x10013000
.equ GPIO_BASE,  0x10012000
.equ GPIO_OUTPUT_EN,  0x08
.equ GPIO_OUTPUT_VAL, 0x0c

#.equ PIN7, 1 # pin no and gpio no are different
.equ PIN1, 1

_start:

# t0 holds GPIO structure address
la t0, GPIO_BASE

# t1 holds the bitmask for pin
li t1, PIN1
li t2, 1
sll t1, t2, t1

# set pin as output
# these can be overwritten
lw t2, GPIO_OUTPUT_EN(t0)
or t3, t1, t2
sw t3, GPIO_OUTPUT_EN(t0)

