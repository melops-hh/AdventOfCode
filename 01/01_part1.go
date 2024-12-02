package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var left, right []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(split[len(split)-1])
		if err != nil {
			panic(err)
		}
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += abs(left[i] - right[i])
	}
	fmt.Println("sum:", sum)
}
