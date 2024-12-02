package main

import (
	"bufio"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseLine(line string) (left int, right int) {
	lineValues := strings.Split(line, "   ")

	left, _ = strconv.Atoi(lineValues[0])
	right, _ = strconv.Atoi(lineValues[1])
	return
}
func loadInput(fileName string) (left []int, right []int, err error) {
	slog.Info("Loading", "filename", fileName)
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		slog.Error("Got err", "err", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		leftValue, rightValue := parseLine(line)
		left = append(left, leftValue)
		right = append(right, rightValue)
	}
	slices.Sort(left)
	slices.Sort(right)

	return
}
func calculatePath(left, right []int) (result int) {

	for index, valueLeft := range left {
		valueRight := right[index]
		diff := valueRight - valueLeft

		if diff < 0 {
			diff *= -1
		}
		result += diff
	}
	return
}
func main() {
	slog.Info("Day solution")
	fileName := "input.txt"
	left, right, err := loadInput(fileName)
	if err != nil {
		return
	}
	result := calculatePath(left, right)
	slog.Info("Path is", "path", result)

}
