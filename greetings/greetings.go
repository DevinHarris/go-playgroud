package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// Declare a variable == var message string
	//var message string = fmt.Sprintf("Hi, %v. Welcome!", name)

	// If no name was given, return an error with a message.

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
