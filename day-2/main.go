package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func checkPath(line string) (ok bool) {
	lineValues := make([]int, 0)
	for _, v := range strings.Split(line, " ") {
		value, _ := strconv.Atoi(v)
		lineValues = append(lineValues, value)
	}
	var isAsc, isDes bool
	for index, value := range lineValues {
		if isAsc && isDes {
			return false
		}
		if index == len(lineValues)-1 {
			break
		}
		diff := value - lineValues[index+1]
		if diff < 0 {
			isDes = true
			diff *= -1
		} else {
			isAsc = true
		}
		if diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}
func countSafePath(fileScanner *bufio.Scanner) (result int) {
	for fileScanner.Scan() {
		line := fileScanner.Text()
		res := checkPath(line)
		if res {
			result += 1
		}
	}
	return
}
func main() {
	slog.Info("Day 2 solution")
	fileName := "input.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		slog.Error("Got err", "err", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	res := countSafePath(fileScanner)
	slog.Info("Part 1", "res", res)
}
