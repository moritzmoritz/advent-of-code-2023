package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	startIndex int
	endIndex   int
	value      int
}

type Symbol struct {
	index     int
	character string
}

type Line struct {
	numbers []*Number
	symbols []*Symbol
}

func main() {
	input, er := os.Open("./input.txt")

	if er != nil {
		panic(er)
	}
	defer input.Close()

	first := partOne(input)
	input.Seek(0, 0)
	second := partTwo(input)

	fmt.Printf("Part one: %d\n", first)
	fmt.Printf("Part two: %d\n", second)
}

func partOne(input *os.File) int {
	sc := bufio.NewScanner(input)

	lines := make([]*Line, 0)

	for sc.Scan() {
		line := sc.Text()
		lineObject := readLine(line)
		lines = append(lines, lineObject)
	}

	total := 0

	for lineIndex, line := range lines {
		for _, number := range line.numbers {
			counted := false
			if lineIndex > 0 {
				// Check line before
				for _, symbol := range lines[lineIndex-1].symbols {
					if symbol.index >= number.startIndex-1 && symbol.index <= number.endIndex+1 {
						total += number.value
						counted = true
					}
				}
			}

			if !counted {
				// check symbols in same line
				for _, symbol := range line.symbols {
					if symbol.index == number.startIndex-1 || symbol.index == number.endIndex+1 {
						total += number.value
						counted = true
					}
				}
			}

			if !counted && lineIndex < len(lines)-1 {
				// Check line after
				for _, symbol := range lines[lineIndex+1].symbols {
					if symbol.index >= number.startIndex-1 && symbol.index <= number.endIndex+1 {
						total += number.value
						counted = true
					}
				}
			}
		}
	}

	return total
}

func partTwo(input *os.File) int {
	sc := bufio.NewScanner(input)

	lines := make([]*Line, 0)

	for sc.Scan() {
		line := sc.Text()
		lineObject := readLine(line)
		lines = append(lines, lineObject)
	}

	total := 0

	for lineIndex, line := range lines {
		for _, symbol := range line.symbols {
			if symbol.character != "*" {
				continue
			}

			fmt.Println("check symbol", symbol.character, symbol.index)
			numbersAround := make([]*Number, 0)
			if lineIndex > 0 {
				for _, number := range lines[lineIndex-1].numbers {
					if number.endIndex >= symbol.index-1 && number.startIndex <= symbol.index+1 {
						numbersAround = append(numbersAround, number)
					}
				}
			}

			// Check same line
			for _, number := range line.numbers {
				if number.endIndex == symbol.index-1 || number.startIndex == symbol.index+1 {
					numbersAround = append(numbersAround, number)
				}
			}

			// Check next line
			if lineIndex < len(lines)-1 {
				for _, number := range lines[lineIndex+1].numbers {
					if number.endIndex >= symbol.index-1 && number.startIndex <= symbol.index+1 {
						numbersAround = append(numbersAround, number)
					}
				}
			}

			if len(numbersAround) == 2 {
				gearRatio := numbersAround[0].value * numbersAround[1].value
				total += gearRatio
			}
		}
	}

	return total
}

func readLine(line string) *Line {
	currentNumber := make([]int, 0)

	lineNumbers := make([]*Number, 0)
	lineSymbols := make([]*Symbol, 0)

	for index, char := range line {
		// Find numbers in each row with their star and end index
		// Find all symbols in each row
		if unicode.IsDigit(rune(char)) {
			number, err := strconv.Atoi(string(char))

			if err != nil {
				continue
			}

			currentNumber = append(currentNumber, number)

			if index == len(line)-1 {
				value := buildInteger(currentNumber)

				number := Number{
					startIndex: index - len(currentNumber) + 1,
					endIndex:   index,
					value:      value,
				}

				currentNumber = make([]int, 0)
				lineNumbers = append(lineNumbers, &number)
			}
		} else {

			if char != '.' {
				symbol := Symbol{
					index:     index,
					character: string(char),
				}

				lineSymbols = append(lineSymbols, &symbol)
			}

			if len(currentNumber) > 0 {
				value := buildInteger(currentNumber)

				number := Number{
					startIndex: index - len(currentNumber),
					endIndex:   index - 1,
					value:      value,
				}

				currentNumber = make([]int, 0)
				lineNumbers = append(lineNumbers, &number)
			}
		}
	}

	lineObject := Line{
		numbers: lineNumbers,
		symbols: lineSymbols,
	}

	return &lineObject
}

func buildInteger(digits []int) int {
	result := 0

	for _, digit := range digits {
		result = result*10 + digit
	}

	return result
}
