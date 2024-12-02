package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sum += isSafe(scanner.Text())
	}
	fmt.Println(sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(s string) int {
	intArr := toInt(s)
	dec := intArr[0]-intArr[1] > 0

	for i := 1; i < len(intArr); i++ {
		diff := intArr[i-1] - intArr[i]
		if dec && diff < 0 {
			return 0
		}
		if !dec && diff > 0 {
			return 0
		}
		if abs(diff) < 1 || abs(diff) > 3 {
			return 0
		}
	}
	return 1
}

func toInt(s string) []int {
	strs := strings.Fields(s)
	arr := make([]int, len(strs))
	for index, strval := range strs {
		ival, _ := strconv.Atoi(strval)
		arr[index] = ival
	}
	return arr

}
