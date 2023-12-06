package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type PastRace struct {
	time           int
	recordDistance int
}

func main() {
	input, er := os.Open("./input.txt")

	if er != nil {
		panic(er)
	}
	defer input.Close()

	first := partOne(input)

	inputTwo, er := os.Open("./input_2.txt")

	if er != nil {
		panic(er)
	}
	defer inputTwo.Close()

	second := partTwo(inputTwo)

	fmt.Printf("Part one: %d\n", first)
	fmt.Printf("Part two: %d\n", second)
}

func partOne(input *os.File) int {
	races := readRaces(input)

	total := 1

	for _, race := range races {
		possibilities := 0
		for millisecond := 1; millisecond < race.time; millisecond++ {
			timeTravelled := race.time - millisecond
			speedPerMillisecond := millisecond
			distanceTravelled := timeTravelled * speedPerMillisecond

			if distanceTravelled > race.recordDistance {
				possibilities++
			}
		}

		total *= possibilities
	}

	return total
}

func partTwo(input *os.File) int {
	races := readRaces(input)

	total := 1

	for _, race := range races {
		possibilities := 0
		for millisecond := 1; millisecond < race.time; millisecond++ {
			timeTravelled := race.time - millisecond
			speedPerMillisecond := millisecond
			distanceTravelled := timeTravelled * speedPerMillisecond

			if distanceTravelled > race.recordDistance {
				possibilities++
			}
		}

		total *= possibilities
	}

	return total
}

func readRaces(input *os.File) []*PastRace {
	sc := bufio.NewScanner(input)

	times := make([]int, 0)
	distances := make([]int, 0)

	lineIndex := 0
	for sc.Scan() {
		line := sc.Text()

		if lineIndex == 0 {
			// Read time
			re := regexp.MustCompile(`\b\d+\b`)
			matches := re.FindAllString(line, -1)

			for _, time := range matches {
				timeNumber, err := strconv.Atoi(time)

				if err != nil {
					continue
				}

				times = append(times, timeNumber)
			}
			lineIndex++
			continue
		}

		// Read distances
		re := regexp.MustCompile(`\b\d+\b`)
		matches := re.FindAllString(line, -1)

		for _, distance := range matches {
			distanceNumber, err := strconv.Atoi(distance)

			if err != nil {
				continue
			}

			distances = append(distances, distanceNumber)
		}
	}

	if len(times) != len(distances) {
		panic("Times and distances do not match")
	}

	races := make([]*PastRace, 0)

	for idx, time := range times {
		distance := distances[idx]

		races = append(races, &PastRace{
			time:           time,
			recordDistance: distance,
		})
	}

	return races
}
