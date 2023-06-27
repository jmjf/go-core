package greeter

import "fmt"

// public
func Welcome(name string) string {
	// Uncomment the lines below to get cover2.out and cover2.html
	if name != "Dave" {
		fmt.Println("I should not be executed in test for Dave -", name)
	}

	return fmt.Sprintf("Welcome to golang, %v!\n", name)
}

// private
func buhbye(name string) string {
	return fmt.Sprintf("Hasta la vista, %v.\n", name)
}
