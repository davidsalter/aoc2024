package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davidsalter/aoc2024/helpers"
)

func Day2Part1() {
	fmt.Println("Day 02 Part 1")

	data := helpers.ReadFileAsStrings("inputs/day02.txt")

	safecount := 0

	for l := 0; l < len(data); l++ {

		tokens := strings.Fields(data[l])

		safe := checkLine(tokens)

		if safe {
			safecount = safecount + 1
		}
	}
	fmt.Println(safecount)
}

func Day2Part2() {
	fmt.Println("Day 02 Part 2")
	data := helpers.ReadFileAsStrings("inputs/day02.txt")

	safecount := 0

	for l := 0; l < len(data); l++ {

		tokens := strings.Fields(data[l])

		maxtokens := len(tokens)
		safe := checkLine(tokens)
		if !safe {
			for i := 0; i < maxtokens; i++ {
				checkTokens := helpers.Remove(tokens, i)
				safe = checkLine(checkTokens)
				if safe {
					break
				}
			}
		}

		if safe {
			safecount = safecount + 1
		}

	}
	fmt.Println(safecount)
}

func checkLine(tokens []string) bool {
	maxtokens := len(tokens) - 1
	increase := 0
	decrease := 0
	for i := 0; i < maxtokens; i++ {
		val, _ := strconv.Atoi(tokens[i])
		valnext, _ := strconv.Atoi(tokens[i+1])
		if valnext >= val {
			if (valnext-val) <= 3 && (valnext-val) > 0 {
				increase = increase + 1
			}
		} else if valnext <= val {
			if (val-valnext) <= 3 && (val-valnext) > 0 {
				decrease = decrease + 1
			}
		}
	}

	if increase == maxtokens && decrease == 0 {
		return true
	} else if increase == 0 && decrease == maxtokens {
		return true
	}
	return false
}
