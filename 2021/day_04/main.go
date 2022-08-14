package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type BoardField struct {
	IsMarked bool
	Number   int
}

type Board [][]BoardField

func generateBoards(rows []string) []Board {
	boards := []Board{}

	currentBoard := Board{}
	for index, row := range rows {
		if index > 0 && index%5 == 0 {
			boards = append(boards, currentBoard)
			currentBoard = Board{}
		}

		row := regexp.MustCompile("[0-9]+").FindAllString(row, -1)
		fields := []BoardField{}
		for _, digit := range row {
			num, _ := strconv.Atoi(digit)
			field := BoardField{
				Number:   num,
				IsMarked: false,
			}
			fields = append(fields, field)
		}
		currentBoard = append(currentBoard, fields)

		if index == len(rows)-1 {
			boards = append(boards, currentBoard)
		}
	}

	return boards
}

func checkWinningBoard(board Board) bool {
	isWinning := false
	reversedBoard := make(Board, len(board[0]))
	for _, row := range board {
		counter := 0
		rowSize := len(row)
		for j, num := range row {
			if num.IsMarked {
				counter++
			}
			reversedBoard[j] = append(reversedBoard[j], num)
		}
		if counter == rowSize {
			isWinning = true
			break
		}
	}

	// check all columns
	if !isWinning {
		for _, column := range reversedBoard {
			counter := 0
			columnSize := len(column)
			for _, num := range column {
				if num.IsMarked {
					counter++
				}
			}
			if counter == columnSize {
				isWinning = true
				break
			}
		}
	}

	return isWinning
}

func findWinningBoard(boards []Board) Board {
	for _, board := range boards {
		isWinningBoard := checkWinningBoard(board)
		if isWinningBoard {
			return board
		}
	}
	return nil
}

func fillInBoards(sequence []string, boards []Board) (Board, int) {
	fmt.Println(sequence)
	var lastNum int
	for _, strNum := range sequence {
		num, _ := strconv.Atoi(strNum)

		winningBoard := findWinningBoard(boards)
		if winningBoard != nil {
			return winningBoard, lastNum
		}

		for i := 0; i < len(boards); i++ {
			for j := 0; j < len(boards[i]); j++ {
				for k := 0; k < len(boards[i][j]); k++ {
					if boards[i][j][k].Number == num {
						boards[i][j][k].IsMarked = true
					}
				}
			}
		}
		lastNum = num
	}

	return nil, lastNum
}

func generateFinalResult(board Board, lastNumber int) int {
	sum := 0
	for _, row := range board {
		for _, field := range row {
			if !field.IsMarked {
				sum += field.Number
			}
		}
	}

	return sum * lastNumber
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputRows := strings.Split(string(content), "\n")

	nonEmptyRows := []string{}

	for _, row := range inputRows {
		if row != "" {
			nonEmptyRows = append(nonEmptyRows, row)
		}
	}

	inputSequence := regexp.MustCompile("[0-9]+").FindAllString(nonEmptyRows[0], -1)

	boards := generateBoards(nonEmptyRows[1:])

	winningBoard, lastNumber := fillInBoards(inputSequence, boards)

	fmt.Println(winningBoard, lastNumber)

	fmt.Println(generateFinalResult(winningBoard, lastNumber))
}
