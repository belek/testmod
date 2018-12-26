package testmod

import "fmt"

// Hi returns greetings
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
