package main

import (
	"testing"
)

func TestCalibrate(t *testing.T) {
	tests := []struct {
		it     string
		input  string
		result int
	}{
		{
			it:     "should return 142",
			input:  "part_one.txt",
			result: 142,
		},
		{
			it:     "should return 281",
			input:  "part_two.txt",
			result: 281,
		},
	}

	for _, test := range tests {
		t.Run(test.it, func(t *testing.T) {
			result := Calibrate(test.input)
			if result != test.result {
				t.Errorf("expected %d, got %d", test.result, result)
			}
		})
	}
}
