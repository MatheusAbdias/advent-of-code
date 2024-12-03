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

func IsSafeReport(report []string) bool {
	isSafeReport := true
	isDecreasing := false

	for i := 0; i < len(report)-1; i++ {
		lv, _ := strconv.Atoi(report[i])
		nextLv, _ := strconv.Atoi(report[i+1])
		diff := nextLv - lv
		absDiff := int(math.Abs(float64(diff)))

		if absDiff == 0 || absDiff > 3 {
			isSafeReport = false
			break
		} else {
			if i == 0 && diff < 0 {
				isDecreasing = true
			}
			if (isDecreasing && diff > 0) || (!isDecreasing && diff < 0) {
				return false
			}
		}
	}

	return isSafeReport
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
		report := strings.Fields(scanner.Text())

		if IsSafeReport(report) {
			safeReports += 1
		}
	}

	fmt.Println("The number of safe reports is: ", safeReports)
}
