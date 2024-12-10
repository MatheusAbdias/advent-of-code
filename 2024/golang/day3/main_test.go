package main

import "testing"

func TestExtractValue(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	expected := 161
	result := ExtractValue(input)
	if result != expected {
		t.Errorf("The expected value is: %v but got: %v", expected, result)
	}
}
