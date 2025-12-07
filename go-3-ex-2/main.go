package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign, "\n")
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	switch zodiacSign {
	case Sagittarius:
		fmt.Printf("23.11-22.12")
	case Scorpius:
		fmt.Printf("24.10-23.11")
	case Libra:
		fmt.Printf("24.09-24.10")
	case Virgo:
		fmt.Printf("24.08-24.09")
	case Leo:
		fmt.Printf("24.07-24.08")
	case Cancer:
		fmt.Printf("22.06-24.07")
	case Gemini:
		fmt.Printf("22.05-22.06")
	case Taurus:
		fmt.Printf("21.04-22.05")
	case Aries:
		fmt.Printf("21.03-21.04")
	case Pisces:
		fmt.Printf("20.02-21.03")
	case Aquarius:
		fmt.Printf("20.02-21.01")
	case Capricornus:
		fmt.Printf("21.01-22.12")
	}
	// TODO: ...and consider a default case.
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
