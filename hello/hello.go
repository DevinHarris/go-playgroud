package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {

	// Set properties of the predefined Loggger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.

	message, err := greetings.Hello("Devin")

	// If an error was returned, print it to the console and exit the program.

	if err != nil {
		log.Fatal(err)
	}

	// If no error print to console

	fmt.Println(message)
}
