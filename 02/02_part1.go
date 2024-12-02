package main

import (
	"bufio"
	"fmt"
	"os"
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
	var reports [][]int
	sumOfSafe := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		var r []int
		isSafe := true
		var isIncreasing bool
		for i := range split {
			a, err := strconv.Atoi(split[i])
			if err != nil {
				panic(err)
			}
			r = append(r, a)
			if i > 0 {
				dif := r[i] - r[i-1]
				absDif := abs(dif)
				// fmt.Printf("i: %d, a: %d, dif: %d, absDif: %d\n", i, a, dif, absDif)
				if absDif < 1 || absDif > 3 {
					// fmt.Println("dif is too big")
					isSafe = false
					break
				}

				if i == 1 {
					if r[i]-r[i-1] < 0 {
						isIncreasing = false
					} else {
						isIncreasing = true
					}
				} else {
					if isIncreasing && dif <= 0 {
						// fmt.Println("inc dif >=0")
						isSafe = false
						break
					}
					if !isIncreasing && dif >= 0 {
						// fmt.Println("dec dif <=0")
						isSafe = false
						break
					}
				}
			}
		}
		if isSafe {
			sumOfSafe++
		}
		reports = append(reports, r)

	}
	fmt.Printf("Sum of Safe: %d\n", sumOfSafe)
}
