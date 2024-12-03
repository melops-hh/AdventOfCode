package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readFile() string {
	data, err := os.ReadFile("input_test2.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseMul(input string) [][]int {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	return regex.FindAllStringSubmatchIndex(input, -1)
}

func parseDont(input string) [][]int {
	regex := regexp.MustCompile(`don't\(\)|do\(\)`)
	matches := regex.FindAllStringSubmatchIndex(input, -1)

	var skips [][]int
	skipStart := 0
	skipping := false

	for i, m := range matches {
		startIndex, endIndex := m[0], m[1]
		match := input[startIndex:endIndex]
		if match == "don't()" && !skipping {
			skipping = true
			skipStart = startIndex
		}
		if match == "do()" && skipping {
			skipping = false
			skips = append(skips, []int{skipStart, endIndex})
		}
		if i == len(matches)-1 && skipping {
			skips = append(skips, []int{skipStart, len(input)})
		}
	}

	return skips
}

func multiply(muls [][]int, donts [][]int, input string) int {
	sum := 0
	for _, m := range muls {
		skip := false
		for _, d := range donts {
			if m[0] >= d[0] && m[1] <= d[1] {
				skip = true
			}
		}

		if !skip {
			a, err := strconv.Atoi(input[m[2]:m[3]])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(input[m[4]:m[5]])
			if err != nil {
				panic(err)
			}
			sum += a * b
		}
	}
	return sum
}

func main() {
	input := readFile()
	muls := parseMul(input)
	donts := parseDont(input)
	sum := multiply(muls, donts, input)
	fmt.Println("Sum:", sum)
}
