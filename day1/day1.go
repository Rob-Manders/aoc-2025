package day1

import (
	"strconv"
)

const MAX = 100

func Rotate(input []string) int {
	counter := 0
	dial := 50

	for _, line := range input {
		operator := line[0]
		amount, _ := strconv.Atoi(line[1:])

		switch operator {
		case 'L':
			dial = rotateLeft(amount, dial)
		case 'R':
			dial = rotateRight(amount, dial)
		}

		if dial == 0 {
			counter++
		}
	}

	return counter
}

func rotateRight(amount, dial int) int {
	v := dial + amount + 1
	for v > MAX {
		v -= MAX
	}

	return v - 1
}

func rotateLeft(amount, dial int) int {
	v := dial - amount + 1
	for v < 0 {
		v += MAX
	}

	return v - 1
}
