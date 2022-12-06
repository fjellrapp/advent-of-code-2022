package dayTwo

import (
	"adventofcode/mats/utils"
	"bufio"
)

const (
	OpponentRock    string = "A"
	OpponentPaper   string = "B"
	OpponentScissor string = "C"
)

const (
	Rock    string = "X"
	Paper   string = "Y"
	Scissor string = "Z"
)

const ()

func DayTwo() {
	partOne()
	partTwo()
}

func partOne() {
	scanner := readFileAndReturnScanner()
	strategyMap := map[string]int{
		"A X": 1 + 3,
		"A Y": 6 + 2,
		"A Z": 0 + 3,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 6 + 1,
		"C Y": 0 + 2,
		"C Z": 3 + 3,
	}

	var pointsSummary int = 0

	for scanner.Scan() {
		stratResult := strategyMap[scanner.Text()]
		pointsSummary += stratResult
	}
}

func partTwo() {
	scanner := readFileAndReturnScanner()
	var pointsSummary int = 0
	strategyMap := map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}

	for scanner.Scan() {
		stratResult := strategyMap[scanner.Text()]
		pointsSummary += stratResult
	}
}

func readFileAndReturnScanner() *bufio.Scanner {
	file := utils.OpenFile("dayTwo/data.txt")

	fileScanner := bufio.NewScanner(file)

	return fileScanner
}
