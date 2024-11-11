# RISCV Board code

- Make sure you have the `riscv-unknown-elf-gcc` toolchain installed
- build the code executable by running `make` within this directory
- flash the executable to the board by using the file browser to 'drag and drop' the main.hex file into the boards flash memory.
- Connect to the serial output of the board using `sudo screen /dev/ttyACM* 115200` or any suitable software to connect to the serial port (putty etc)

# 
- Check 202411112.csv for a trace from `main.s` running on the board.


## TODO
- Write a script to generate the assembly files with random operands and registers
- automate flashing these files onto the board after the last one has finished running and the trace has been saved
