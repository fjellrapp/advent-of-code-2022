package dayThree

import (
	"adventofcode/mats/utils"
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

var letterRangeLower = make(map[string]int)
var letterRangeUpper = make(map[string]int)
var currentLowerIndex int = 1
var currentUpperIndex int = 27

func DayThree() {
	constructLetterRange()
	partOne()
	partTwo()
}

func constructLetterRange() {
	for i := 'a'; i <= 'z'; i++ {

		letterRangeLower[string(i)] = currentLowerIndex
		currentLowerIndex++
	}
	for i := 'A'; i <= 'Z'; i++ {
		letterRangeUpper[string(i)] = currentUpperIndex
		currentUpperIndex++

	}
}

func partOne() {

	scanner := readFileAndReturnScanner()

	commonItemsMap := make(map[string]string)
	for scanner.Scan() {
		var stringLength = len(scanner.Text())
		compartmentOne := scanner.Text()[0 : stringLength/2]
		compartmentTwo := scanner.Text()[stringLength/2:]

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

func partTwo() {
	scanner := readFileAndReturnScanner()

	groups := constructGroups(scanner)

	var sum int = 0
	for _, group := range groups {
		for key, pri := range letterRangeLower {
			if strings.Contains(group[0], string(key)) && strings.Contains(group[1], key) && strings.Contains(group[2], string(key)) {
				sum += pri
				break
			}
		}
		for key, pri := range letterRangeUpper {
			if strings.Contains(group[0], string(key)) && strings.Contains(group[1], string(key)) && strings.Contains(group[2], string(key)) {
				sum += pri
				break
			}
		}
	}

	fmt.Println(sum)

}

func constructGroups(scanner *bufio.Scanner) map[int][]string {
	groups := make(map[int][]string)

	var MAX_ITERATION int = 3
	var CURRENT int = 0
	var GROUP_INDEX int = 0
	for scanner.Scan() {
		if CURRENT == MAX_ITERATION {
			CURRENT = 0
			GROUP_INDEX += 1
		}

		groups[GROUP_INDEX] = append(groups[GROUP_INDEX], scanner.Text())
		CURRENT++
	}
	return groups
}
func readFileAndReturnScanner() *bufio.Scanner {
	file := utils.OpenFile("dayThree/data.txt")

	fileScanner := bufio.NewScanner(file)

	return fileScanner
}
