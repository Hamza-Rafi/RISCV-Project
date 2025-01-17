package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// var instruction = "add"
var instructions = []string{"add", "sub", "sll", "sltu", "xor", "srl", "sra", "or", "and"}
var type2 = []string{"addi", "slti", "sltiu", "xori", "ori", "andi", "slli", "srli", "srai", "lui", "auipc", "lb", "lh", "lw", "lbu", "lhu"}
var type3 = []string{"sb", "sh", "sw"}
var type4 = []string{"beq", "bne", "blt", "bge", "bltu", "bgeu"}
var type5 = []string{"lui", "auipc"}
var type6 = []string{"jal", "jalr"}

func getRandomRegister(registers []string) string {
	rn := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rn)

	randRegister := registers[r.Intn(len(registers))]

	registerIndex := getIndexOfItem(registers, randRegister)
	registers = append(registers[:registerIndex], registers[registerIndex+1:]...)

	return randRegister
}

func getRandomOperand() string {
	// max signed integer value
	max := 2147483647
	rn := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rn)

	randNum := r.Intn(max)

	return "0x" + strconv.FormatInt(int64(randNum), 16)
}

func getIndexOfItem(slice []string, item string) int {
	index := -1
	for i, elem := range slice {
		if item == elem {
			index = i
			break
		}
	}
	return index
}

func main() {
	// fileBeginning, err := os.ReadFile("fileBeginning.s")
	// open files
	loop, err := os.Open("loop.s")
	if err != nil {
		panic(err)
	}

	sections := make([][]string, 0)
	scanner := bufio.NewScanner(loop)
	section := []string{}

	// populate sections array
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			section = append(section, line)
		} else {
			sections = append(sections, section)
			section = []string{}
		}
	}

	// write file Beginning
	fileBeginning, err := os.ReadFile("fileBeginning.s")
	if err != nil {
		panic(err)
	}

	for _, instruction := range instructions {
		// create new file
		newFile, err := os.Create("./out/" + instruction + ".s")
		defer newFile.Close()

		if err != nil {
			panic(err)
		}
		writer := bufio.NewWriter(newFile)

		// fmt.Print(string(fileBeginning))
		writer.Write(fileBeginning)
		writer.Flush()

		// loop 100 times
		for count := 0; count < 100; count++ {
			var registers = []string{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10", "x11", "x12", "x13", "x14", "x15", "x16", "x17", "x18", "x19", "x20", "x21", "x22", "x23", "x24", "x25", "x26", "x27", "x28", "x29", "x30"}

			var r1, r2, r3, v1, v2 string

			writer.WriteString("#-------------------" + strconv.Itoa(count) + "---------------------\n")

			for i, section := range sections {
				// generate registers and values only for the first section
				if i == 0 {
					r1, r2 = getRandomRegister(registers), getRandomRegister(registers)
					v1, v2 = getRandomOperand(), getRandomOperand()
				}

				for _, line := range section {
					// prerequisite section
					if i == 0 {
						// replace placeholders
						line = strings.Replace(line, "{r1}", r1, 1)
						line = strings.Replace(line, "{r2}", r2, 1)
						line = strings.Replace(line, "{v1}", v1, 1)
						line = strings.Replace(line, "{v2}", v2, 1)
					} else if i == 2 {
						// instruction section
						// generate r3
						r3 = getRandomRegister(registers)

						// replace placeholders
						line = strings.Replace(line, "{instruction}", instruction, 1)
						line = strings.Replace(line, "{r1}", r1, 1)
						line = strings.Replace(line, "{r2}", r2, 1)
						line = strings.Replace(line, "{r3}", r3, 1)
					}

					writer.WriteString(line + "\n")
				}
			}
			writer.WriteString("#-----------------------------------------\n\n")
			writer.Flush()

			// fmt.Println(sections)
		}
	}
}
