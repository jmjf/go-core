package greeter

import "testing"
import "fmt"

// Time to create tests (white box)

func TestWelcome(t *testing.T) {
	name := "Dave"
	expected := "Welcome to golangg, Dave!\n"
	got := Welcome(name)

	if got != expected {
		t.Errorf("Expected: %v", expected)
		t.Fatalf("Got: %v", got)
	}

	fmt.Println("I'm here!")
}

func TestBuhbye(t *testing.T) {
	t.Errorf("Whoops")
	fmt.Println("And I'm here")
}
