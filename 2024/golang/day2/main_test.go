package main

import (
	"testing"
)

func TestIsSafeReport(t *testing.T) {
	type TestCase struct {
		description   string
		input         []int
		expectedError bool
	}
	for _, scenario := range []TestCase{
		{
			description:   "Safe without removing any level.",
			input:         []int{7, 6, 4, 2, 1},
			expectedError: false,
		},
		{
			description:   "Unsafe regardless of which level is removed.",
			input:         []int{1, 2, 7, 8, 9},
			expectedError: true,
		},
		{
			description:   "Unsafe regardless of which level is removed.",
			input:         []int{9, 7, 6, 2, 1},
			expectedError: true,
		},
		{
			description:   "Safe by removing the second level, 3.",
			input:         []int{1, 3, 2, 4, 5},
			expectedError: false,
		},
		{
			description:   "Safe by removing the third level, 4.",
			input:         []int{8, 6, 4, 4, 1},
			expectedError: false,
		},
		{
			description:   "Safe without removing any level.",
			input:         []int{1, 3, 6, 7, 9},
			expectedError: false,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			isSafe := TryRemoveToBeSafe(scenario.input)

			if isSafe && scenario.expectedError || !isSafe && !scenario.expectedError {
				t.Error(scenario.description)
			}
		})
	}
}
