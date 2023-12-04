package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToint(s string) string {
	substrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	it := 1
	for _, substring := range substrings {
		for {
			idx := strings.Index(s, substring)
			if idx == -1 {
				break
			}
			substringLength := len(substring)
			s = s[:idx+(substringLength/2)] + strconv.Itoa(it) + s[idx+(substringLength/2):]
		}
		it++
	}
	return s
}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		first := 0
		last := 0
		isFirst := true
		lineTransformed := stringToint(line)

		for _, char := range lineTransformed {
			if char >= 48 && char <= 57 {
				if isFirst {
					first = int(char) - 48
					isFirst = false
				} else {
					last = int(char) - 48
				}
			}
		}
		if last == 0 {
			last = first
		}
		sum += (first*10 + last)
	}
	fmt.Println(sum)
}
