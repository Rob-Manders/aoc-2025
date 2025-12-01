package main

import (
	"aoc/day1"
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput("./inputs/day1.txt")
	password := day1.Rotate(input)

	fmt.Println(password)
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}
