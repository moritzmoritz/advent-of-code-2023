package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, er := os.Open("./input.txt")

	if er != nil {
		panic(er)
	}
	defer input.Close()

	first := partOne(input, 12, 13, 14)
	input.Seek(0, 0)
	second := partTwo(input)

	fmt.Printf("Part one: %d\n", first)
	fmt.Printf("Part two: %d\n", second)
}

func partOne(input *os.File, redCubes int, greenCubes int, blueCubes int) int {
	sc := bufio.NewScanner(input)

	total := 0

	for sc.Scan() {
		line := sc.Text()

		splittedLine := strings.Split(line, ":")

		gameString := strings.Split(splittedLine[0], " ")
		id, err := strconv.Atoi(gameString[1])

		if err != nil {
			continue
		}

		minRed := 0
		minGreen := 0
		minBlue := 0

		rounds := strings.Split(splittedLine[1], ";")

		for _, round := range rounds {
			cubes := strings.Split(round, ",")

			for _, amount := range cubes {
				amount = strings.Trim(amount, " ")

				splittedAmount := strings.Split(amount, " ")
				value, err := strconv.Atoi(splittedAmount[0])

				if err != nil {
					continue
				}

				if strings.Contains(splittedAmount[1], "red") {
					if minRed < value {
						minRed = value
					}
				} else if strings.Contains(splittedAmount[1], "green") {
					if minGreen < value {
						minGreen = value
					}
				} else if strings.Contains(splittedAmount[1], "blue") {
					if minBlue < value {
						minBlue = value
					}
				}
			}

		}

		if minRed <= redCubes && minGreen <= greenCubes && minBlue <= blueCubes {
			total += id
		}
	}

	return total
}

func partTwo(input *os.File) int {
	sc := bufio.NewScanner(input)

	total := 0

	for sc.Scan() {
		line := sc.Text()

		splittedLine := strings.Split(line, ":")

		redValue := 0
		greenValue := 0
		blueValue := 0

		rounds := strings.Split(splittedLine[1], ";")

		for _, round := range rounds {
			cubes := strings.Split(round, ",")

			for _, amount := range cubes {
				amount = strings.Trim(amount, " ")

				splittedAmount := strings.Split(amount, " ")
				value, err := strconv.Atoi(splittedAmount[0])

				if err != nil {
					continue
				}

				if strings.Contains(splittedAmount[1], "red") {
					if redValue < value {
						redValue = value
					}
				} else if strings.Contains(splittedAmount[1], "green") {
					if greenValue < value {
						greenValue = value
					}
				} else if strings.Contains(splittedAmount[1], "blue") {
					if blueValue < value {
						blueValue = value
					}
				}
			}

		}

		total += redValue * greenValue * blueValue
	}

	return total
}
