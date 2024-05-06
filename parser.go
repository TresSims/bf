package main

import (
	"bufio"
	"os"
	"strings"
)

const VALID_TOKENS string = "><+-,.[]"

func ParseFile(filepath string) (output string, err error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return ParseString(string(content))
}

func ParseString(input string) (output string, err error) {
	var instructions []rune
	for _, v := range []rune(input) {
		if strings.IndexRune(VALID_TOKENS, v) != -1 {
			instructions = append(instructions, v)
		}
	}

	return Interpret(instructions)
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

func Interpret(instructions []rune) (output string, err error) {

	reader := bufio.NewReader(os.Stdin)

	cell_pointer := 0
	instruction_pointer := 0
	var cells []int

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
			output += string(rune(cells[cell_pointer]))
		case ',':
			r, _, err := reader.ReadRune()
			if err != nil {
				return "", err
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

	return
}
