package day2

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	idRanges := strings.Split(input[0], ",")

	total := 0
	for _, idRange := range idRanges {
		total += validateIDRange(idRange)
	}

	return total
}

func validateIDRange(idRange string) (subTotal int) {
	subTotal = 0

	bounds := strings.Split(idRange, "-")

	lower, _ := strconv.Atoi(bounds[0])
	upper, _ := strconv.Atoi(bounds[1])

	for i := lower; i <= upper; i++ {
		if !isValidID(i) {
			subTotal += i
		}
	}

	return subTotal
}

func isValidID(input int) bool {
	idString := strconv.Itoa(input)

	part1 := idString[0 : len(idString)/2]
	part2 := idString[len(idString)/2:]

	return part1 != part2
}
