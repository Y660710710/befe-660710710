package main

import (
	"fmt"
)

// var email string = "Thanayod.21@gmail.com"

func main() {
	// var name string = "Thanayod"
	var age int = 20

	email := "Thanayod.21@gmail.com"
	gpa := 4.00

	firstname, lastname := "Thanayod", "Rattanasureeroj"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)
}
