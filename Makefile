all: as link hex

as:
	riscv64-unknown-elf-as -march=rv32imac -mno-arch-attr main.s -o main.o

link: 
	riscv64-unknown-elf-ld -T linker.ld main.o -b elf32-littleriscv -o main.elf

hex:
	riscv64-unknown-elf-objcopy -O ihex main.elf main.hex



