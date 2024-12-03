package main

import (
	"aoc24/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var left, right []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		l := utils.ToInt(split[0])
		r := utils.ToInt(split[len(split)-1])
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += utils.Abs(left[i] - right[i])
	}
	fmt.Println("sum:", sum)
}
