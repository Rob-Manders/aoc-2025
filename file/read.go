package file

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filename string) []string {
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
