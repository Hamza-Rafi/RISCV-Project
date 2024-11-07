# RISCV Board code

- Make sure you have the `riscv-unknown-elf-gcc` toolchain installed
- build the code executable by running `make` within this directory
- flash the executable to the board by using the file browser to 'drag and drop' the main.hex file into the boards flash memory.
- Connect to the serial output of the board using `sudo screen /dev/ttyACM* 115200` or any suitable software to connect to the serial port (putty etc)
