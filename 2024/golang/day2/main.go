package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func TryRemoveToBeSafe(report []int) bool {
	for index := 0; index < len(report); index++ {
		current := make([]int, len(report))
		copy(current, report)

		ok := IsSafeReport(append(current[:index], current[index+1:]...))
		if ok {
			return true
		}
	}

	return false
}

func ConvertReport(report []string) []int {
	var result []int
	for _, value := range report {
		vInt, _ := strconv.Atoi(value)
		result = append(result, vInt)
	}

	return result
}

func IsSafeReport(report []int) bool {
	isDecreasing := false

	for i := 0; i < len(report)-1; i++ {
		lv := report[i]
		nextLv := report[i+1]
		diff := nextLv - lv
		absDiff := int(math.Abs(float64(diff)))

		if absDiff == 0 || absDiff > 3 {
			return false
		}
		if i == 0 && diff < 0 {
			isDecreasing = true
		}
		if (isDecreasing && diff > 0) || (!isDecreasing && diff < 0) {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		result := strings.Fields(scanner.Text())
		report := ConvertReport(result)
		if TryRemoveToBeSafe(report) {
			safeReports += 1
		}
	}

	fmt.Println("The number of safe reports is: ", safeReports)
}
