package dayOne

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panicMessage := fmt.Sprintf("Fails: %f", e)
		panic(panicMessage)
	}
}
func readAndReturnFile() []byte {
	dat, err := os.ReadFile("dayOne/data.txt")
	check(err)
	return dat
}

func getAccumulatedCalories() []int {
	dat := readAndReturnFile()
	fileString := string(dat)

	stringsSlice := strings.Split(fileString, "\n\n")

	accumulatedCalories := make([]int, 0)

	for _, i := range stringsSlice {
		var sum int = 0
		splitGroup := strings.Split(i, "\n")
		for _, j := range splitGroup {
			if j != "" {
				value, err := strconv.Atoi(j)
				if err != nil {
					panic(err)
				}
				sum += value
			}

		}
		accumulatedCalories = append(accumulatedCalories, sum)
	}

	sort.Ints(accumulatedCalories)
	return accumulatedCalories

}

func findLargest(calories []int) int {
	largest := calories[len(calories)-1]
	return largest
}

func DoWork() {
	calories := getAccumulatedCalories()
	greatest := findLargest(calories)
	fmt.Print(greatest)
}
