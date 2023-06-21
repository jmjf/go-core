package greeter

import "fmt"

// public
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to golang, %v!\n", name)
}

// private
func buhbye(name string) string {
	return fmt.Sprintf("Hasta la vista, %v.\n", name)
}
