package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileAsStrings(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()
	return lines
}

func Abs(number int) int {
	if number > 0 {
		return number
	}
	return -1 * number
}

func Remove(slice []string, index int) []string {
	newSlice := make([]string, 0, len(slice))
	for i, v := range slice {
		if i != index {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
