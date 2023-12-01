package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Digit struct {
	index int
	value int
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

	total := 0

	for sc.Scan() {
		line := sc.Text()
		digits := make([]int, 0)

		characters := strings.Split(line, "")
		for _, c := range characters {
			digit, err := strconv.Atoi(c)

			if err != nil {
				continue
			}

			digits = append(digits, digit)
		}

		first := digits[0]
		last := digits[len(digits)-1]

		value, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

		if err != nil {
			continue
		}

		total += value
	}

	return total
}

func partTwo(input *os.File) int {
	characterWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sc := bufio.NewScanner(input)

	total := 0
	for sc.Scan() {
		digits := make([]Digit, 0)

		line := sc.Text()

		for wordIndex, word := range characterWords {
			indexes := make([]int, 0)

			// Find all indexes of the word
			for i := 0; ; {
				index := strings.Index(line[i:], word)
				if index == -1 {
					break
				}

				indexes = append(indexes, i+index)
				i += index + 1
			}

			for _, index := range indexes {
				digits = append(digits, Digit{index, wordIndex + 1})
			}
		}

		characters := strings.Split(line, "")
		for index, c := range characters {
			digit, err := strconv.Atoi(c)

			if err != nil {
				continue
			}

			digits = append(digits, Digit{index, digit})
		}

		if len(digits) == 0 {
			continue
		}

		first := digits[0]
		last := digits[len(digits)-1]

		for _, value := range digits {
			if value.index < first.index {
				first = value
			} else if value.index > last.index {
				last = value
			}
		}

		value, err := strconv.Atoi(fmt.Sprintf("%d%d", first.value, last.value))

		if err != nil {
			continue
		}

		total += value
	}

	return total
}
