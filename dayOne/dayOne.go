package dayOne

import (
	"adventofcode/mats/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getAccumulatedCalories() []int {
	dat := utils.ReadAndReturnFile("dayOne/data.txt")
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
	fmt.Println(greatest)
}
