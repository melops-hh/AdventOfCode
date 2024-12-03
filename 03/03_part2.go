package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Wrong answers:
// 78965138 to high
// 146499494
// 153469856 to high

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		sum += handleText(scanner.Text())
	}
	fmt.Println("sum:", sum)
}

func handleText(s string) int {
	fmt.Println("\nhandle new line\n")
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don\\'t\\(\\)")
	match := r.FindAllString(s, -1)
	sum := 0
	enabled := true
	for _, mul := range match {
		newDo, s := multiply(mul, enabled)
		fmt.Printf("multiply %s, do: %t, sum: %d, newDo: %t newSum: %d\n", mul, enabled, sum, newDo, s)
		if enabled {
			sum += s
		}
		enabled = newDo
	}
	return sum
}

func multiply(s string, enabled bool) (bool, int) {
	if s == "do()" {
		return true, 0
	} else if s == "don't()" || !enabled {
		return false, 0
	}

	s = s[4 : len(s)-1]
	split := strings.Split(s, ",")
	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])
	return true, a * b

}
