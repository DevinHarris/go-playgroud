package greetings

import "fmt"

func Hello(name string) string {
	// Declare a variable == var message string
	//var message string = fmt.Sprintf("Hi, %v. Welcome!", name)

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
