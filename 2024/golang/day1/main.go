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

func insertSorted(toInsert []int, value int) []int {
	left, right := 0, len(toInsert)

	for left < right {
		mid := (left + right) / 2

		if toInsert[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}

	toInsert = append(toInsert, 0)
	copy(toInsert[left+1:], toInsert[left:])
	toInsert[left] = value
	return toInsert
}

func main() {
	var leftColumn, rightColumn []int

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
	for scanner.Scan() {
		places := strings.Fields(scanner.Text())
		if len(places) != 2 {
			log.Fatalf("Invalid line format: %s", scanner.Text())
		}

		leftValue, err := strconv.Atoi(places[0])
		if err != nil {
			log.Fatalf("Error parsing left value: %v", err)
		}
		leftColumn = insertSorted(leftColumn, leftValue)

		rightValue, err := strconv.Atoi(places[1])
		if err != nil {
			log.Fatalf("Error parsing right value: %v", err)
		}
		rightColumn = insertSorted(rightColumn, rightValue)
	}

	var sum float64
	for i := 0; i < len(leftColumn); i++ {
		diff := math.Abs(float64(leftColumn[i] - rightColumn[i]))
		sum += diff
	}

	fmt.Printf("The total distance between left and right columns is: %.2f\n", sum)
}
