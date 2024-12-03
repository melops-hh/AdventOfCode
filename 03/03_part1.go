package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		sum += handleText(scanner.Text())
	}
	fmt.Println(sum)
}

func handleText(s string) int {
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	match := r.FindAllString(s, -1)
	sum := 0
	for _, mul := range match {
		sum += multiply(mul)
	}
	return sum
}

func multiply(s string) int {
	s = s[4 : len(s)-1]
	split := strings.Split(s, ",")
	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])
	return a * b
}
