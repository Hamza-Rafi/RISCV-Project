OUTPUT_ARCH(riscv)
OUTPUT_FORMAT("elf32-littleriscv")
ENTRY(_start)

SECTIONS{
	. = 0x20010000;
	.text : { *(.text) }

	.rodata : { *(.rodata) }

	. = 0x80000000;
	.data  : { *(.data) }	
}
