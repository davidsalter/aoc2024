package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/davidsalter/aoc2024/helpers"
)

func Day1Part1() {
	fmt.Println("Day 01 Part 1")
	left := []int{}
	right := []int{}
	data := helpers.ReadFileAsStrings("inputs/day01.txt")
	for _, line := range data {
		tokens := strings.Fields(line)
		l, _ := strconv.Atoi(strings.TrimSpace(tokens[0]))
		r, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		distance = distance + helpers.Abs(left[i]-right[i])
	}

	fmt.Println(distance)
}

func Day1Part2() {
	fmt.Println("Day 01 Part 2")
	left := []int{}
	right := []int{}
	data := helpers.ReadFileAsStrings("inputs/day01.txt")
	for _, line := range data {
		tokens := strings.Fields(line)
		l, _ := strconv.Atoi(strings.TrimSpace(tokens[0]))
		r, _ := strconv.Atoi(strings.TrimSpace(tokens[1]))
		left = append(left, l)
		right = append(right, r)
	}

	similarity := 0
	for i := 0; i < len(left); i++ {
		similarity += (left[i] * appear(left[i], right))
	}

	fmt.Println(similarity)
}

func appear(l int, r []int) int {
	count := 0
	for v := range r {
		if r[v] == l {
			count = count + 1
		}
	}
	return count
}
