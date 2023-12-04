package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	number         int
	winningNumbers []int
	playerNumbers  []int
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

	cards := make([]*Card, 0)

	lineIndex := 0
	for sc.Scan() {
		line := sc.Text()
		card := readCard(line, lineIndex, 5)
		cards = append(cards, card)
		lineIndex++
	}

	total := 0

	for _, card := range cards {
		cardNumbers := make([]int, 0)

		for _, number := range card.playerNumbers {
			for _, winningNumber := range card.winningNumbers {
				if number == winningNumber {
					cardNumbers = append(cardNumbers, number)
				}
			}
		}

		if len(cardNumbers) > 0 {
			cardTotal := 0

			for idx, _ := range cardNumbers {
				if idx == 0 {
					cardTotal = 1
				} else {
					cardTotal *= 2
				}
			}
			total += cardTotal
		}
	}

	return total
}

func partTwo(input *os.File) int {
	sc := bufio.NewScanner(input)

	cards := make([]*Card, 0)

	lineIndex := 0
	for sc.Scan() {
		line := sc.Text()
		card := readCard(line, lineIndex, 10)
		cards = append(cards, card)
		lineIndex++
	}

	total := len(cards)

	copies := make([]*Card, 0)

	for idx, card := range cards {
		cardNumbers := make([]int, 0)

		for _, number := range card.playerNumbers {
			for _, winningNumber := range card.winningNumbers {
				if number == winningNumber {
					cardNumbers = append(cardNumbers, number)
				}
			}
		}

		if len(cardNumbers) > 0 {
			copies = append(copies, cards[idx+1:idx+1+len(cardNumbers)]...)
			total += len(cardNumbers)
		}
	}

	for len(copies) > 0 {
		newCopies := make([]*Card, 0)

		for _, card := range copies {
			cardNumbers := make([]int, 0)

			for _, number := range card.playerNumbers {
				for _, winningNumber := range card.winningNumbers {
					if number == winningNumber {
						cardNumbers = append(cardNumbers, number)
					}
				}
			}

			if len(cardNumbers) > 0 {
				idx := card.number - 1
				newCopies = append(newCopies, cards[idx+1:idx+1+len(cardNumbers)]...)
			}
		}

		total += len(newCopies)
		copies = newCopies
	}

	return total
}

func readCard(line string, idx int, winningNumbersCount int) *Card {
	splittedCard := strings.Split(line, ":")
	splittedNumbers := strings.Split(splittedCard[1], "|")

	winningNumbers := make([]int, 0)
	playerNumbers := make([]int, 0)

	for _, input := range splittedNumbers {
		numbers := strings.Split(input, " ")

		for _, number := range numbers {
			intNumber, err := strconv.Atoi(number)

			if err != nil {
				continue
			}

			if len(winningNumbers) < winningNumbersCount {
				winningNumbers = append(winningNumbers, intNumber)
			} else {
				playerNumbers = append(playerNumbers, intNumber)
			}
		}
	}

	card := Card{
		number:         idx + 1,
		winningNumbers: winningNumbers,
		playerNumbers:  playerNumbers,
	}

	return &card
}
