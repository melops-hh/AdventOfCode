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

func checkArr(intArr []int) bool {
	inc, dec := false, false

	for i := 1; i < len(intArr); i++ {
		diff := intArr[i] - intArr[i-1]
		if diff > 0 {
			inc = true
		} else if diff < 0 {
			dec = true
		} else {
			return false
		}
		if dec && inc {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
	}

	return true

}

func isDelSafe(intArr []int, index int) bool {
	newArr := make([]int, len(intArr))
	copy(newArr, intArr)

	if index == len(newArr)-1 {
		newArr = newArr[:index]
	} else {
		newArr = append(newArr[:index], newArr[index+1:]...)
	}
	return checkArr(newArr)
}

func isSafe(s string) int {
	intArr := toInt(s)
	if checkArr(intArr) {
		return 1
	} else {
		for i := 0; i < len(intArr); i++ {
			if isDelSafe(intArr, i) {
				return 1
			}
		}
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
