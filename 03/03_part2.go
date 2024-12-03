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
	fmt.Println("sum:", sum)
}

func handleText(s string) int {
	fmt.Println("\nhandle new line\n")
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don\\'t\\(\\)")
	match := r.FindAllString(s, -1)
	sum := 0
	lastInstruction := 0
	for _, mul := range match {
		i, s := multiply(mul, lastInstruction)
		fmt.Printf("multiply %s, lastInstruction: %d, sum: %d, newSum: %d\n", mul, lastInstruction, sum, s)
		lastInstruction = i
		if i == 0 || i == 1 {
			sum += s
		}
	}
	return sum
}

func multiply(s string, lastI int) (int, int) {
	instruction := strings.Split(s, "(")[0]
	if instruction == "do" {
		return 1, 0
	} else if instruction == "don't" {
		return 2, 0
	} else if instruction == "mul" {
		s := s[4 : len(s)-1]
		split := strings.Split(s, ",")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		return lastI, a * b
	} else {
		panic("Invalid instruction")
	}

}
