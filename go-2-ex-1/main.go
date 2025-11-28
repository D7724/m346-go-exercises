package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
	// TODO: add fields
}

// TODO: declare a structure for birth date
type BirthDate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {

	var birthDate = BirthDate{
		dayOfBirth:   26,
		monthOfBirth: 12,
		yearOfBirth:  2008,
	}
	var fullName = FullName{"John", "Doe"}

	var me = Profile{
		// TODO: set name and birth date information
		FullName:         fullName,
		BirthDate:        birthDate,
		NumberOfSiblings: 4,        // TODO: adjust
		ZodiacSign:       '\u264B', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
