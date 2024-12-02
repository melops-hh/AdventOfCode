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

func checkArr(intArr []int) (bool, int) {
	dec := intArr[0]-intArr[1] > 0

	for i := 1; i < len(intArr); i++ {
		diff := intArr[i-1] - intArr[i]
		if dec && diff < 0 {
			return false, i
		}
		if !dec && diff > 0 {
			return false, i
		}
		if abs(diff) < 1 || abs(diff) > 3 {
			return false, i
		}
	}

	return true, 0

}

func isSafe(s string) int {
	intArr := toInt(s)
	safe, i := checkArr(intArr)
	if safe {
		return 1
	}
	intArr2 := append(intArr[:i], intArr[i+1:]...)
	safe1, _ := checkArr(intArr2)
	if safe1 {
		return 1
	}
	return 0
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
