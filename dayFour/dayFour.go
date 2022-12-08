package dayFour

import (
	"adventofcode/mats/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var sum_fully_contains int = 0
var sum_overlap int = 0

func DayFour() {
	scanner := readFileAndReturnScanner()

	for scanner.Scan() {
		splitRange := strings.Split(scanner.Text(), ",")
		formattedRangeUpper := formatRange(splitRange[0])
		formattedRangeLower := formatRange(splitRange[1])

		fully_contains(formattedRangeUpper, formattedRangeLower)
	}

	fmt.Print("overlap", sum_overlap)

}

func fully_contains(upper, lower []int) {
	var a int = upper[0]
	var b int = upper[1]
	var c int = lower[0]
	var d int = lower[1]

	if c <= a && b <= d || a <= c && d <= b {
		sum_fully_contains += 1
	}

	var maxVal = max(a, c)
	var minVal = min(b, d)

	if maxVal <= minVal {
		sum_overlap += 1
	}
}

func min(one, two int) int {
	if one > two {
		return two
	} else {
		return one
	}
}

func max(one, two int) int {
	if one > two {
		return one
	} else {
		return two
	}
}

func formatRange(splitRange string) []int {
	out := make([]int, 0)
	replaceDash := strings.Split(splitRange, "-")
	indexOne, _ := strconv.Atoi(replaceDash[0])
	indexTwo, _ := strconv.Atoi(replaceDash[1])
	return append(out, indexOne, indexTwo)
}

func readFileAndReturnScanner() *bufio.Scanner {
	file := utils.OpenFile("dayFour/data.txt")

	fileScanner := bufio.NewScanner(file)

	return fileScanner
}
