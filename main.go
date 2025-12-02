package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/file"
	"fmt"
	"os"
	"strconv"
)

type day struct {
	Part1 func([]string) int
	Part2 func([]string) int
}

var days = map[int]day{
	1: {day1.Part1, day1.Part2},
	2: {day2.Part1, nil},
}

func main() {
	dayArg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	partArg, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	dayFuncs, ok := days[dayArg]
	if !ok {
		fmt.Printf("Day %d does not exist", dayArg)
		return
	}

	var partFunc func([]string) int
	switch partArg {
	case 1:
		partFunc = dayFuncs.Part1
	case 2:
		partFunc = dayFuncs.Part2
	default:
		fmt.Printf("Invalid art arg: %d", partArg)
		return
	}

	if partFunc == nil {
		fmt.Println("Part function is nil")
		return
	}

	testFile := fmt.Sprintf("./inputs/day%d.txt", dayArg)
	if len(testFile) == 0 {
		fmt.Println("Test input is empty")
	}

	input := file.Read(testFile)
	output := partFunc(input)

	fmt.Println(output)
}
