package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkKnown(known map[int]int, i int) (bool, int) {
	for n, m := range known {
		if n == i {
			return true, m
		}
	}
	return false, 0

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

	known := map[int]int{}
	sum := 0
	for i := range left {
		isKnown, knownVal := checkKnown(known, left[i])
		if isKnown {
			sum += knownVal
		} else {
			appearances := 0
			for j := range right {
				if left[i] == right[j] {
					appearances++
				}
			}
			sumI := appearances * left[i]
			known[left[i]] = sumI
			sum += sumI
		}

	}

	fmt.Println("sum:", sum)
}
