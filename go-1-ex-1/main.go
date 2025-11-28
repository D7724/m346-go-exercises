package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Ryan"
	lastName := "Trut"
	var dayOfBirth = 26
	var monthOfBirth int = 12
	var yearOfBirth int = 2008
	var numberOfSiblings int = 4
	var heightInMeters float32 = 3.0
	var zodiacSign rune = '\u264B'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
