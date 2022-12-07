package dayThree

import (
	"adventofcode/mats/utils"
	"bufio"
	"fmt"
	"unicode"
)

func DayThree() {
	partOne()
}

func partOne() {

	scanner := readFileAndReturnScanner()
	letterRangeLower := make(map[string]int)
	letterRangeUpper := make(map[string]int)
	var currentLowerIndex int = 1
	var currentUpperIndex int = 27
	for i := 'a'; i <= 'z'; i++ {

		letterRangeLower[string(i)] = currentLowerIndex
		currentLowerIndex++
	}
	for i := 'A'; i <= 'Z'; i++ {
		letterRangeUpper[string(i)] = currentUpperIndex
		currentUpperIndex++

	}
	commonItemsMap := make(map[string]string)
	for scanner.Scan() {
		var stringLength = len(scanner.Text())
		compartmentOne := scanner.Text()[0 : stringLength/2]
		compartmentTwo := scanner.Text()[stringLength/2:]

		fmt.Println(compartmentOne, compartmentTwo)
		for _, one := range compartmentOne {
			for _, two := range compartmentTwo {
				if unicode.IsUpper(one) && unicode.IsUpper(two) {
					if string(one) == string(two) {
						commonItemsMap[scanner.Text()] = string(one)
						break
					}
				}
				if unicode.IsLower(one) && unicode.IsLower(two) {
					if string(one) == string(two) {
						commonItemsMap[scanner.Text()] = string(one)
						break
					}
				}
			}
		}

	}

	var sum int = 0
	for _, mapItem := range commonItemsMap {
		if letterRangeLower[mapItem] > 0 {
			sum += letterRangeLower[mapItem]
		}
		if letterRangeUpper[mapItem] > 0 {
			sum += letterRangeUpper[mapItem]
		}
	}

	fmt.Println(sum)
}
func readFileAndReturnScanner() *bufio.Scanner {
	file := utils.OpenFile("dayThree/data.txt")

	fileScanner := bufio.NewScanner(file)

	return fileScanner
}
