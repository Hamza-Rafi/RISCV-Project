.section .text
.globl _start

.equ UART_BASE,  0x10013000
.equ GPIO_BASE,  0x10012000
.equ GPIO_OUTPUT_EN,  0x08
.equ GPIO_OUTPUT_VAL, 0x0c

#.equ PIN7, 1 # pin no and gpio no are different
.equ PIN1, 1

_start:
	la t0, GPIO_BASE

	# crete bitmask for pin 
	# bitmask stored in t1
	li t1, PIN1
	li t2, 1
	sll t1, t2, t1

	# set pin as output
	lw t2, GPIO_OUTPUT_EN(t0)
	or t3, t1, t2
	sw t3, GPIO_OUTPUT_EN(t0)

	# prerequisite for instruction
	li t4, 5
	li t5, 6

	# set pin high
	lw t2, GPIO_OUTPUT_VAL(t0)
	or t3, t1, t2
	sw t3, GPIO_OUTPUT_VAL(t0)

	# run instruction
	add t4, t4, t5
	
	# set pin low
	not t1, t1
	lw t2, GPIO_OUTPUT_VAL(t0)
	and t3, t1, t2
	sw t3, GPIO_OUTPUT_VAL(t0)

# THIS CODE IS FOR DEBUGGING PURPOSES.
# WILL NOT INCLUDE WHEN RUNNING THE TESTS
# ---------------------------------
	la a0, text
	jal print_string

print_string:
	li t0, UART_BASE
print_char:
	lb a1, (a0)
	beqz a1, end_print
wait:
	lw t1, (t0)
	bltz t1, wait

	sw a1, (t0)
	addi a0, a0, 1
	j print_char

end_print:
	ret

.section .rodata
high: .string "high\r\n"
low: .string "low\r\n"
text:
	.string "finish\r\n"
# ---------------------------------
