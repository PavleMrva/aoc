package main

import (
	"testing"
)

func TestCalculateGameIDs(t *testing.T) {
	tests := []struct {
		it                   string
		input                string
		config               *GameConfig
		expectedSumOfGameIDs int
		expectedSumOfPowers  int
	}{
		{
			it:    "should return 8",
			input: "part_one.txt",
			config: &GameConfig{
				Blue:  14,
				Green: 13,
				Red:   12,
			},
			expectedSumOfGameIDs: 8,
			expectedSumOfPowers:  2286,
		},
	}

	for _, test := range tests {
		t.Run(test.it, func(t *testing.T) {
			sumOfGameIDs, sumOfPowers := CalculateGameIDs(test.config, test.input)
			if sumOfGameIDs != test.expectedSumOfGameIDs {
				t.Errorf("expected sumOfGameIDs %d, got %d", test.expectedSumOfGameIDs, sumOfGameIDs)
			}

			if sumOfPowers != test.expectedSumOfPowers {
				t.Errorf("expected sumOfGameIDs %d, got %d", test.expectedSumOfPowers, sumOfPowers)
			}
		})
	}
}
