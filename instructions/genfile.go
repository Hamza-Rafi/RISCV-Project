package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var instruction = "add"

func getRandomRegister(registers []string, shouldPop bool) string {
	rn := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rn)

	randRegister := registers[r.Intn(len(registers))]

	// for prerequisite values, dont overwrite the registers
	if shouldPop {
		registerIndex := getIndexOfItem(registers, randRegister)
		registers = append(registers[:registerIndex], registers[registerIndex+1:]...)
	}

	return randRegister
}

func getRandomNumber(max int) int {
	rn := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rn)

	randNum := r.Intn(max)

	return randNum
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

	// create new file
	newFile, err := os.Create("test.s")
	defer newFile.Close()

	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(newFile)

	for count := 0; count < 100; count++ {
		var registers = []string{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10", "x11", "x12", "x13", "x14", "x15", "x16", "x17", "x18", "x19", "x20", "x21", "x22", "x23", "x24", "x25", "x26", "x27", "x28", "x29", "x30"}
		var r1, r2, r3, r4 string

		writer.WriteString("-------------------" + strconv.Itoa(count) + "---------------------\n")

		for i, section := range sections {

			if i == 0 {
				r1, r2 = getRandomRegister(registers, true), getRandomRegister(registers, true)
			}

			r3, r4 = getRandomRegister(registers, false), getRandomRegister(registers, false)

			// makes sure r4 and r4 are not the same
			for r3 == r4 {
				r4 = getRandomRegister(registers, false)
			}

			for _, line := range section {
				// if line is comment, ignore
				if line[0] == '#' {
					fmt.Println(line)
					writer.WriteString(line + "\n")
					continue
				}

				// (2^32 -1) / 2
				maxNum := 2147483647

				// prerequisite instructions
				if i == 0 {
					line = strings.Replace(line, "{r1}", r1, 1)
					line = strings.Replace(line, "{r2}", r2, 1)

					randNum := getRandomNumber(maxNum)
					line = strings.Replace(line, "{v1}", "0x"+strconv.FormatInt(int64(randNum), 16), 1)

					randNum = getRandomNumber(maxNum)
					line = strings.Replace(line, "{v2}", "0x"+strconv.FormatInt(int64(randNum), 16), 1)

				} else if i == 2 {
					line = strings.Replace(line, "{instruction}", instruction, 1)
					line = strings.Replace(line, "{r1}", r1, 1)
					line = strings.Replace(line, "{r2}", r2, 1)

					line = strings.Replace(line, "{r3}", getRandomRegister(registers, false), 1)

				} else {
					// replace r3 and r4 with random registers
					line = strings.Replace(line, "{r3}", r3, 1)
					line = strings.Replace(line, "{r4}", r4, 1)

					// replace r1 and r2 with operand registers
					line = strings.Replace(line, "{r1}", r1, 1)
					line = strings.Replace(line, "{r2}", r2, 1)

				}
				fmt.Println(line)

				// write line to file
				writer.WriteString(line + "\n")

			}

		}
		writer.WriteString("-----------------------------------------")
		writer.Flush()

		fmt.Println(sections)
	}
}
