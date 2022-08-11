package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne(measurements []int) {
	counter := 0
	for index, _ := range measurements {
		if index != 0 && measurements[index] > measurements[index-1] {
			counter++
		}
	}

	fmt.Println(counter)
}

func partTwo(measurements []int) {
	currSum := 0
	previousSum := 0
	counter := 0
	for index, _ := range measurements {
		if index != 0 && index+2 < len(measurements) {
			previousSum = currSum
			currSum = measurements[index] + measurements[index+1] + measurements[index+2]

			if currSum > previousSum {
				counter++
			}
		}
	}
	fmt.Println(counter)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	measurements := []int{}
	for _, measurement := range strings.Split(string(content), "\n") {
		measurementNumber, _ := strconv.Atoi(measurement)
		measurements = append(measurements, measurementNumber)
	}

	partOne(measurements)
	partTwo(measurements)
}
