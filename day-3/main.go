package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInput(fileName string) (string, error) {
	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body), nil
}
func sumMuls(text string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(text, -1)
	total := 0
	for _, match := range matches {
		x, err1 := strconv.Atoi(match[1])
		y, err2 := strconv.Atoi(match[2])
		if err1 == nil && err2 == nil {
			total += x * y
		}
	}

	return total
}

func sumWithConditionals(memory string) int {
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	allInstructions := regexp.MustCompile(
		`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`).FindAllString(memory, -1)

	enabled := true
	total := 0

	for _, instruction := range allInstructions {
		if doRe.MatchString(instruction) {
			enabled = true
		} else if dontRe.MatchString(instruction) {
			enabled = false
		} else if mulMatch := mulRe.FindStringSubmatch(instruction); mulMatch != nil {
			if enabled {
				x, _ := strconv.Atoi(mulMatch[1])
				y, _ := strconv.Atoi(mulMatch[2])
				total += x * y
			}
		}
	}

	return total
}

func main() {
	fileName := "input.txt"
	text, err := readInput(fileName)
	if err != nil {
		return
	}
	fmt.Println("Part 1", sumMuls(text))
	fmt.Println("Part 2", sumWithConditionals(text))
}
