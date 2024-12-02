package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"
)

func checkPath(line string) (ok bool) {
	lineValues := make([]int, 0)
	for _, v := range strings.Split(line, " ") {
		value, _ := strconv.Atoi(v)
		lineValues = append(lineValues, value)
	}
	return isSafe(lineValues)
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

func isSafe(lineValues []int) bool {
	if len(lineValues) < 2 {
		return true
	}
	isAsc := true
	isDes := true
	for i := 0; i < len(lineValues)-1; i++ {
		diff := lineValues[i+1] - lineValues[i]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if diff > 0 {
			isDes = false
		} else if diff < 0 {
			isAsc = false
		}

		if !isAsc && !isDes {
			return false
		}
	}
	return true
}
func checkPathWithTolerance(line string) (ok bool) {
	lineValues := make([]int, 0)
	for _, v := range strings.Split(line, " ") {
		value, _ := strconv.Atoi(v)
		lineValues = append(lineValues, value)
	}
	if isSafe(lineValues) {
		return true
	}
	// TODO think about optimization
	for i := 0; i < len(lineValues); i++ {
		modified := append([]int{}, lineValues[:i]...)
		modified = append(modified, lineValues[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}

	return false

}
func countSafeReportsOnePass(fileScanner *bufio.Scanner) (result int) {
	for fileScanner.Scan() {
		line := fileScanner.Text()
		res := checkPathWithTolerance(line)
		// fmt.Println(line, res)
		if res {
			result += 1
		}
	}
	return
}
func getPartOne(wg *sync.WaitGroup) {
	defer wg.Done()
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
func getPartTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	fileName := "input.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		slog.Error("Got err", "err", err)
		return
	}
	fileScanner := bufio.NewScanner(file)
	res := countSafeReportsOnePass(fileScanner)
	slog.Info("Part 2", "res", res)
}

func main() {
	slog.Info("Day 2 solution")
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go getPartOne(wg)
	go getPartTwo(wg)
	wg.Wait()
}
