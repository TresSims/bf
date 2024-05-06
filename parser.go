package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const VALID_TOKENS string = "><+-,.[]"

func ParseFile(filepath string) {
	content, error := os.ReadFile(filepath)
	if error != nil {
		fmt.Println("File " + filepath + " not found. Exiting.")
		return
	}

	ParseString(string(content))
}

func ParseString(input string) {
	var instructions []rune
	for _, v := range []rune(input) {
		if strings.IndexRune(VALID_TOKENS, v) != -1 {
			instructions = append(instructions, v)
		}
	}

	Interpret(instructions)
}

func matching_brace(instructions []rune, instruction_pointer int, find_right bool) (location int) {
	travel_further := 0
	if find_right {
		location = instruction_pointer + 1
	} else {
		location = instruction_pointer - 1
	}

	if find_right {
		for location < len(instructions) {
			switch instructions[location] {
			case '[':
				travel_further++
			case ']':
				if travel_further == 0 {
					return
				}
				travel_further--
			}

			location++
		}
	} else {
		for location >= 0 {
			switch instructions[location] {
			case ']':
				travel_further++
			case '[':
				if travel_further == 0 {
					return
				}
				travel_further--
			}

			location--
		}
	}

	panic("Matching brace not found!")
}

func Interpret(instructions []rune) {

	reader := bufio.NewReader(os.Stdin)

	cell_pointer := 0
	instruction_pointer := 0
	var cells []int

	output := ""

	for instruction_pointer < len(instructions) {

		if cell_pointer >= len(cells) {
			cells = append(cells, 0)
		}

		switch instructions[instruction_pointer] {
		case '>':
			cell_pointer++
		case '<':
			cell_pointer--
		case '+':
			cells[cell_pointer]++
		case '-':
			cells[cell_pointer]--
		case '.':
			output += string(cells[cell_pointer])
		case ',':
			r, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("ERROR READING CHARACTER!")
				return
			}
			cells[cell_pointer] = int(r)
		case '[':
			if cells[cell_pointer] == 0 {
				instruction_pointer = matching_brace(instructions, instruction_pointer, true)
			}
		case ']':
			if cells[cell_pointer] != 0 {
				instruction_pointer = matching_brace(instructions, instruction_pointer, false)
			}
		}

		instruction_pointer++
	}

	fmt.Println(output)
	return
}
