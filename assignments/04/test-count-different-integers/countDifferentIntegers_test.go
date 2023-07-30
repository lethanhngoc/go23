package main

import (
	"testing"
)

func TestCountDifferentIntegers(t *testing.T) {
	result := countDifferentIntegers("a123bc34d8ef34")
	expected := 4
	if result != expected {
		t.Errorf("Expected: %d, but got %d", expected, result)
	}
}
