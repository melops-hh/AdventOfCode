package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		if val, exists := known[left[i]]; exists {
			sum += val
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
