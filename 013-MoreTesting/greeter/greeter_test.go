package greeter_test

import (
	"moreTesting/greeter"
	"testing"
)

// Time to create tests (white box)

func TestWelcome(t *testing.T) {
	name := "Dave"
	// If I remove the \n, the test fails
	// I'm not using fmt.Sprintf here because testing by
	// duplicating function logic is bad practice.
	expected := "Welcome to golang, Dave!\n"
	got := greeter.Welcome(name)

	if got != expected {
		// Using %q wraps the string in quotes and includes escaped characters,
		// which gets "<blah blah blah>\n", which is more useful for comparing.
		// If I use %v, I get no quotes or \n, which makes it easy to miss
		// differences in trailing (and other) whitespace.
		t.Errorf("Expected: %q", expected)
		t.Errorf("Got: %q", got)
	}
}

// func TestBuhbye(t *testing.T) {
// 	name := "Dave"
// 	expected := "Hasta la vista, Dave.\n"
// 	got := greeter.buhbye(name)

// 	if got != expected {
// 		t.Errorf("Expected: %v", expected)
// 		t.Errorf("Got: %v", got)
// 	}
// }
