package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var validDigits map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Calibrate(input string) int {
	content, err := os.ReadFile(input)
	if err != nil {
		return 0
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	var result int

	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}

		num1 := 0
		num2 := 0
		minIdx := math.MaxInt32
		maxIdx := math.MinInt32

		for k, v := range validDigits {
			if !strings.Contains(lines[i], k) {
				continue
			}

			firstIdx := strings.Index(lines[i], k)
			lastIdx := strings.LastIndex(lines[i], k)

			if firstIdx < minIdx {
				minIdx = firstIdx
				num1 = v
			}

			if lastIdx > maxIdx {
				maxIdx = lastIdx
				num2 = v
			}
		}

		val := fmt.Sprintf("%d%d", num1, num2)

		valNum, err := strconv.Atoi(val)
		if err != nil {
			log.Println(err)
			continue
		}

		result += valNum
	}

	return result
}

func main() {
	result := Calibrate("real_input.txt")
	log.Println(result)
}
