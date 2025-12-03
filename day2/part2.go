package day2

import (
	"strconv"
	"strings"
)

func Part2(input []string) int {
	idRanges := strings.Split(input[0], ",")

	total := 0
	for _, idRange := range idRanges {
		total += validateIDRangeWithRepeatingPatterns(idRange)
	}

	return total
}

func validateIDRangeWithRepeatingPatterns(idRange string) (subTotal int) {
	subTotal = 0

	bounds := strings.Split(idRange, "-")

	lower, _ := strconv.Atoi(bounds[0])
	upper, _ := strconv.Atoi(bounds[1])

	for i := lower; i <= upper; i++ {
		if !isValidIDWithRepeatingPatterns(i) {
			subTotal += i
		}
	}

	return subTotal
}

func isValidIDWithRepeatingPatterns(input int) bool {
	id := strconv.Itoa(input)

	divisors := getDivisors(id)
	for _, d := range divisors {
		segments := splitID(id, d)

		if hasRepeatingPattern(segments) {
			return false
		}
	}

	return true
}

func hasRepeatingPattern(segments []string) bool {
	for i := 1; i < len(segments); i++ {
		if segments[i-1] != segments[i] {
			return false
		}
	}

	return true
}

func splitID(id string, divisor int) []string {
	segmentLength := len(id) / divisor
	segments := []string{id[:segmentLength]}

	var i int
	for i = segmentLength; i < len(id)-segmentLength; i += segmentLength {
		end := segmentLength + i
		segment := id[i:end]

		segments = append(segments, segment)
	}

	segments = append(segments, id[i:])

	return segments
}

func getDivisors(id string) []int {
	var divisors []int

	for i := 2; i <= len(id); i++ {
		if len(id)%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}
