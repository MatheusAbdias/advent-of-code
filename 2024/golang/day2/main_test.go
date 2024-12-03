package main

import (
	"testing"
)

func TestIsSafeReport(t *testing.T) {
	type TestCase struct {
		description   string
		input         []string
		expectedError bool
	}

	for _, scenario := range []TestCase{
		{
			description:   "Safe because the levels are all decreasing by 1 or 2.",
			input:         []string{"7", "6", "4", "2", "1"},
			expectedError: false,
		},
		{
			description:   "Unsafe because 2 7 is an increase of 5",
			input:         []string{"1", "2", "7", "8", "9"},
			expectedError: true,
		},
		{
			description:   " Unsafe because 6 2 is a decrease of 4.",
			input:         []string{"9", "7", "6", "2", "1"},
			expectedError: true,
		},

		{
			description:   "Unsafe because 1 3 is increasing but 3 2 is decreasing",
			input:         []string{"1", "3", "2", "4", "5"},
			expectedError: true,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			isSafe := IsSafeReport(scenario.input)

			if isSafe && scenario.expectedError || !isSafe && !scenario.expectedError {
				t.Error(scenario.description)
			}
		})
	}
}
