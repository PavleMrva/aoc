package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type GameConfig struct {
	Blue  int
	Green int
	Red   int
}

const (
	ColorBlue  = "blue"
	ColorGreen = "green"
	ColorRed   = "red"
)

func getGameID(line string) int {
	strArr := strings.Split(line, " ")
	if len(strArr) < 2 {
		return 0
	}

	gameIDSubstr := strings.Split(line, " ")[1]
	gameIDStr := strings.TrimSuffix(gameIDSubstr, ":")

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		return 0
	}

	return gameID
}

func getSets(line string) []string {
	gameSubstr := strings.Split(line, ":")[1]
	gameSubstr = strings.TrimPrefix(gameSubstr, " ")

	sets := strings.Split(gameSubstr, ";")

	return sets
}

func getNoByColorPerSet(set []string, color string) int {
	for i := range set {
		set[i] = strings.TrimPrefix(set[i], " ")
		colorInfo := strings.Split(set[i], " ")

		currNo := colorInfo[0]
		currColor := colorInfo[1]

		if currColor == color {
			no, err := strconv.Atoi(currNo)
			if err != nil {
				return 0
			}

			return no
		}
	}

	return 0
}

func CalculateGameIDs(cfg *GameConfig, input string) (totalPossibleGameIDs int, totalSumOfPower int) {
	content, err := os.ReadFile(input)
	if err != nil {
		return 0, 0
	}

	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	for i := range lines {
		isPossible := true

		gameID := getGameID(lines[i])
		if gameID == 0 {
			continue
		}

		sets := getSets(lines[i])

		maxBlue := 0
		maxGreen := 0
		maxRed := 0

		for j := range sets {
			subsets := strings.Split(sets[j], ",")

			blueNo := getNoByColorPerSet(subsets, ColorBlue)
			greenNo := getNoByColorPerSet(subsets, ColorGreen)
			redNo := getNoByColorPerSet(subsets, ColorRed)

			if blueNo > maxBlue {
				maxBlue = blueNo
			}

			if greenNo > maxGreen {
				maxGreen = greenNo
			}

			if redNo > maxRed {
				maxRed = redNo
			}

			if blueNo > cfg.Blue || greenNo > cfg.Green || redNo > cfg.Red {
				log.Printf("Game ID %d is not possible", gameID)
				isPossible = false
			}
		}

		if isPossible {
			totalPossibleGameIDs += gameID
		}

		totalSumOfPower += maxBlue * maxGreen * maxRed
	}

	return totalPossibleGameIDs, totalSumOfPower
}

func main() {
	cfg := &GameConfig{
		Blue:  14,
		Green: 13,
		Red:   12,
	}

	sumOfGameIDs, sumOfPowers := CalculateGameIDs(cfg, "real_input.txt")
	println(sumOfGameIDs)
	println(sumOfPowers)
}
