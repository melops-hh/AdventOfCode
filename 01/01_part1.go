package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func quickSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}
		}
	}
	return a
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

	left = quicksort(left)
	right = quicksort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		if right[i]-left[i] < 0 {
			sum += left[i] - right[i]
		} else {

			sum += right[i] - left[i]
		}
	}
	fmt.Println("sum:", sum)
}
