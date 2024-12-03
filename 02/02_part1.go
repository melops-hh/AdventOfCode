package main

import (
	"aoc24/utils"
	"bufio"
	"fmt"
	"os"
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
		if utils.Abs(diff) < 1 || utils.Abs(diff) > 3 {
			return 0
		}
	}
	return 1
}

func toInt(s string) []int {
	strs := strings.Fields(s)
	arr := make([]int, len(strs))
	for index, strval := range strs {
		i := utils.ToInt(strval)
		arr[index] = i
	}
	return arr

}
