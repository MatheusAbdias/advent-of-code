package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ExtractValue(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)

	matches := regex.FindAllStringSubmatch(input, -1)

	resultado := 0
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])

		if n1 <= 999 && n2 <= 999 {
			resultado += n1 * n2
		}
	}

	return resultado
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
	result := 0
	for scanner.Scan() {
		result += ExtractValue(scanner.Text())
	}

	fmt.Printf("results of the multiplications %v \n", result)
}
