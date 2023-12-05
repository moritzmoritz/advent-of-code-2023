package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Seed struct {
	number int
}

type CategoryMap struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type Category struct {
	title string
	items []*CategoryMap
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

	// Read seeds
	seeds := make([]*Seed, 0)
	categories := make([]*Category, 0)

	lineIndex := 0
	items := make([]*CategoryMap, 0)

	for sc.Scan() {
		line := sc.Text()

		if lineIndex == 0 {
			// Read seeds
			seeds = readSeeds(line, false)
			lineIndex++
			continue
		}

		// Read category map
		categoryMap := readCategoryMap(line)
		items = append(items, categoryMap)
	}

	category := &Category{
		title: "test",
		items: make([]*CategoryMap, 0),
	}

	for idx, item := range items {
		if item != nil {
			category.items = append(category.items, item)

			if idx == len(items)-1 {
				categories = append(categories, category)
			}
		} else {
			if len(category.items) > 0 {
				categories = append(categories, category)
			}

			category = &Category{
				title: "test",
				items: make([]*CategoryMap, 0),
			}
		}
	}

	lowestLocation := -1
	for _, seed := range seeds {
		source := seed.number

		for _, category := range categories {
			for _, item := range category.items {
				start := item.sourceRangeStart
				end := item.sourceRangeStart + item.rangeLength

				if source >= start && source <= end {
					offset := source - item.sourceRangeStart
					source = item.destinationRangeStart + offset
					break
				}
			}
		}

		if lowestLocation == -1 || source < lowestLocation {
			lowestLocation = source
		}
	}

	return lowestLocation
}

func partTwo(input *os.File) int {
	sc := bufio.NewScanner(input)

	// Read seeds
	seeds := make([]*Seed, 0)
	categories := make([]*Category, 0)

	lineIndex := 0
	items := make([]*CategoryMap, 0)

	for sc.Scan() {
		line := sc.Text()

		if lineIndex == 0 {
			// Read seeds
			seeds = readSeeds(line, true)
			lineIndex++
			continue
		}

		// Read category map
		categoryMap := readCategoryMap(line)
		items = append(items, categoryMap)
	}

	category := &Category{
		title: "test",
		items: make([]*CategoryMap, 0),
	}

	for idx, item := range items {
		if item != nil {
			category.items = append(category.items, item)

			if idx == len(items)-1 {
				categories = append(categories, category)
			}
		} else {
			if len(category.items) > 0 {
				categories = append(categories, category)
			}

			category = &Category{
				title: "test",
				items: make([]*CategoryMap, 0),
			}
		}
	}

	lowestLocation := -1
	for _, seed := range seeds {
		source := seed.number

		for _, category := range categories {
			for _, item := range category.items {
				start := item.sourceRangeStart
				end := item.sourceRangeStart + item.rangeLength

				if source >= start && source <= end {
					offset := source - item.sourceRangeStart
					source = item.destinationRangeStart + offset
					break
				}
			}
		}

		if lowestLocation == -1 || source < lowestLocation {
			lowestLocation = source
		}
	}

	return lowestLocation
}

func readSeeds(input string, seedRange bool) []*Seed {
	if seedRange {
		re := regexp.MustCompile(`\b\d+\b`)
		matches := re.FindAllString(input, -1)

		seeds := make([]int, 0)

		for _, match := range matches {
			seedNumber, err := strconv.Atoi(match)

			if err != nil {
				continue
			}

			seeds = append(seeds, seedNumber)
		}

		var chunked [][]int

		chunkSize := 2

		for i := 0; i < len(seeds); i += chunkSize {
			end := i + chunkSize

			if end > len(seeds) {
				end = len(seeds)
			}

			chunked = append(chunked, seeds[i:end])
		}

		values := make([]*Seed, 0)

		for _, chunk := range chunked {
			for i := 0; i <= chunk[1]; i++ {
				seed := &Seed{
					number: chunk[0] + i,
				}
				values = append(values, seed)
			}
		}

		return values
	}

	seeds := make([]*Seed, 0)

	re := regexp.MustCompile(`\b\d+\b`)
	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			continue
		}

		seed := &Seed{
			number: num,
		}
		seeds = append(seeds, seed)
	}

	return seeds
}

func readCategoryMap(input string) *CategoryMap {
	re := regexp.MustCompile(`\b\d+\b`)
	matches := re.FindAllString(input, -1)
	if len(matches) != 3 {
		return nil
	}

	destinationRangeStart, err := strconv.Atoi(matches[0])

	if err != nil {
		return nil
	}

	sourceRangeStart, err := strconv.Atoi(matches[1])

	if err != nil {
		return nil
	}

	rangeLength, err := strconv.Atoi(matches[2])

	if err != nil {
		return nil
	}

	categoryMap := &CategoryMap{
		destinationRangeStart: destinationRangeStart,
		sourceRangeStart:      sourceRangeStart,
		rangeLength:           rangeLength,
	}

	return categoryMap
}
