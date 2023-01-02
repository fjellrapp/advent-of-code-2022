package dayFive

import (
	"adventofcode/mats/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var cargo map[int][]string = make(map[int][]string, 0)

func DayFive() {

	startingStructure := readAndReturnFile("dayFive/example.txt")
	instructions := readAndReturnFile("dayFive/data.txt")

	createStartingStructureFromFile(startingStructure)

	if len(cargo) != 0 {
		fmt.Println(cargo)
		parseInstructions(instructions)
	}

}

func parseInstructions(instructionsScanner *bufio.Scanner) {
	for instructionsScanner.Scan() {
		instructions := strings.Split(instructionsScanner.Text(), " ")

		move, errMove := strconv.Atoi(instructions[0])
		from, errFrom := strconv.Atoi(instructions[1])
		to, errTo := strconv.Atoi(instructions[2])

		utils.Check(errFrom)
		utils.Check(errMove)
		utils.Check(errTo)

		var out = cargo[from]
		var itemToMove []string

		if len(out) > 1 && len(out) <= move {
			var moveIndex int = len(out) - move
			fmt.Print("index", moveIndex)
			itemToMove = append(itemToMove, out[moveIndex])
		} else {
			itemToMove = append(itemToMove, out[move])
		}

		var in = cargo[to]

		cargo[to] = append(cargo[to], itemToMove...)

		fmt.Println("out, in, move", out, in, move, itemToMove)

	}

	fmt.Println(cargo)

}

func removeElementInRange(s []string, elements int) ([]string, error) {
	if elements > len(s) || elements < 0 {
		return nil, fmt.Errorf("index out of range")
	}

	s = s[:elements]

	fmt.Print("going out", s)

	return s, nil
}

func createStartingStructureFromFile(scanner *bufio.Scanner) {
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "\n")
		for _, element := range splitLine {
			var splitChar = strings.Split(element, " ")
			mapIndex, err := strconv.Atoi(splitChar[0])
			if err != nil {
				panic(err)
			}
			for i, char := range splitChar {
				if i > 0 {
					cargo[mapIndex] = append(cargo[mapIndex], char)
				}
			}
		}
	}
}
func readAndReturnFile(filePath string) *bufio.Scanner {
	file := utils.OpenFile(filePath)

	fileScanner := bufio.NewScanner(file)

	return fileScanner
}
