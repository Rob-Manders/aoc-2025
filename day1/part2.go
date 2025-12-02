package day1

import (
	"strconv"
)

func Part2(input []string) int {
	counter := 0
	dial := 50

	for _, line := range input {
		operator := line[0]
		amount, _ := strconv.Atoi(line[1:])

		switch operator {
		case 'L':
			newDial, zeroes := rotateLeftAndCountZeroes(amount, dial)
			dial = newDial
			counter += zeroes
		case 'R':
			newDial, zeroes := rotateRightAndCountZeroes(amount, dial)
			dial = newDial
			counter += zeroes
		}

		if dial == 0 {
			counter++
		}
	}

	return counter
}

func rotateRightAndCountZeroes(amount, dial int) (newDial, zeroes int) {
	zeroes = 0
	dial += 1

	newDial = dial + amount
	for newDial > MAX {
		zeroes += 1
		newDial -= MAX
	}

	if newDial == 1 {
		zeroes -= 1
	}

	if newDial == 0 {
		newDial = 1
	}

	return newDial - 1, zeroes
}

func rotateLeftAndCountZeroes(amount, dial int) (newDial, zeroes int) {
	zeroes = 0
	dial += 1

	newDial = dial - amount
	for newDial < 0 {
		zeroes += 1
		newDial += MAX
	}

	if newDial == 1 {
		zeroes -= 1
	}

	if newDial == 0 {
		newDial = 99
	}

	return newDial - 1, zeroes
}
