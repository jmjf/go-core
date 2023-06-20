package main_test

import "testing"

// TestMultiplication passes
func TestMultiplication(t *testing.T) {
	name := "When calculating 3 * 9, it returns 27"
	got := 3 * 9
	expected := 27

	if got != expected {
		t.Errorf("TEST: %v\n\tExpected: %v; Got: %v", name, expected, got)
	}

}

// TestDivision fails
func TestDivision(t *testing.T) {
	name := "When calculating 27 / 3, it returns 9"
	got := 27 / 4
	expected := 9

	if got != expected {
		t.Errorf("\nTEST: %v\nExpected: %v; Got: %v", name, expected, got)
	}
}
