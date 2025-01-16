all: as link hex

as:
	riscv64-unknown-elf-as -march=rv32imac -mno-arch-attr $(FILE).s -g -o $(FILE).o

link: 
	riscv64-unknown-elf-ld -T linker.ld $(FILE).o -b elf32-littleriscv -o $(FILE).elf

hex:
	riscv64-unknown-elf-objcopy -O ihex $(FILE).elf $(FILE).hex


