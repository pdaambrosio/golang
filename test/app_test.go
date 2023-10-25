package main

import "testing"

// Function name must start with Test and take a pointer to testing.T as parameter (t *testing.T)
func TestAdd(t *testing.T) {
	// arrange (set up variables) and expected result
	x, y := 1, 2
	expected := 3

	// act (call the function) and get the actual result
	got := Add(x, y)

	// assert (compare actual result with expected result) and print error message if not equal
	if got != expected {
		t.Errorf("Failed Add(%d, %d) = %d; expected %d", x, y, got, expected)
	}
}
