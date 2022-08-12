package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne(content string) {
	binaryInput := strings.Split(content, "\n")
	binaryNumbers := make([][]string, len(binaryInput))

	for i, binarySequence := range binaryInput {
		binaryNumbers[i] = strings.Split(binarySequence, "")
	}

	sequences := make([]string, len(binaryInput[0]))
	for _, binaryNumber := range binaryNumbers {
		for j, digit := range binaryNumber {
			sequences[j] += digit
		}
	}

	gammaRate := ""
	epsilonRate := ""
	for _, sequence := range sequences {
		oneCount := strings.Count(sequence, "1")
		zeroCount := strings.Count(sequence, "0")

		if oneCount > zeroCount {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)
	fmt.Println(gamma * epsilon)
}

func deriveBinNums(rating string, previousBinNums [][]string, sequences []string, counter int) string {
	sequence := sequences[counter]
	oneCount := strings.Count(sequence, "1")
	zeroCount := strings.Count(sequence, "0")

	digits := strings.Split(sequence, "")
	newBinNums := [][]string{}
	for i, digit := range digits {
		if rating == "oxygen" {
			if oneCount >= zeroCount && digit == "1" {
				newBinNums = append(newBinNums, previousBinNums[i])
			}
			if zeroCount > oneCount && digit == "0" {
				newBinNums = append(newBinNums, previousBinNums[i])
			}
		} else {
			if oneCount >= zeroCount && digit == "0" {
				newBinNums = append(newBinNums, previousBinNums[i])
			}
			if zeroCount > oneCount && digit == "1" {
				newBinNums = append(newBinNums, previousBinNums[i])
			}
		}
	}

	if len(newBinNums) == 1 {
		ratingStr := ""
		for _, digit := range newBinNums[0] {
			ratingStr += digit
		}
		return ratingStr
	}

	newSequences := make([]string, len(newBinNums[0]))
	for _, binaryNumber := range newBinNums {
		for j, digit := range binaryNumber {
			newSequences[j] += digit
		}
	}
	counter++
	return deriveBinNums(rating, newBinNums, newSequences, counter)
}

func partTwo(content string) {
	binaryInput := strings.Split(content, "\n")
	binaryNumbers := make([][]string, len(binaryInput))

	for i, binarySequence := range binaryInput {
		binaryNumbers[i] = strings.Split(binarySequence, "")
	}

	sequences := make([]string, len(binaryNumbers[0]))
	for _, binaryNumber := range binaryNumbers {
		for j, digit := range binaryNumber {
			sequences[j] += digit
		}
	}

	oxygenChan := make(chan int64)
	co2Chan := make(chan int64)
	go func() {
		oxygenRatingBin := deriveBinNums("oxygen", binaryNumbers, sequences, 0)
		oxygenRating, _ := strconv.ParseInt(oxygenRatingBin, 2, 64)
		oxygenChan <- oxygenRating
	}()

	go func() {
		co2RatingBin := deriveBinNums("co2", binaryNumbers, sequences, 0)
		co2Rating, _ := strconv.ParseInt(co2RatingBin, 2, 64)
		co2Chan <- co2Rating
	}()

	oxygen := <-oxygenChan
	co2 := <-co2Chan

	fmt.Println("Oxygen rating:")
	fmt.Println(oxygen)

	fmt.Println("CO2 rating:")
	fmt.Println(co2)

	fmt.Println("Final result:")
	fmt.Println(oxygen * co2)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*** PART ONE ***")
	partOne(string(content))
	fmt.Println("*** PART TWO ***")
	partTwo(string(content))
}
