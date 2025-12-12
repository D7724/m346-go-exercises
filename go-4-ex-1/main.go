package main

import (
	"fmt"
	"os"
)

// TODO: implement the function computeGrade

func computeGrade(gotPoints float32, maxPoints float32) float32 {
	if gotPoints > maxPoints || gotPoints < 0 || maxPoints < 0 {
		fmt.Fprintln(os.Stderr, "Not valid input")
		return 0.0
	}
	return (gotPoints/maxPoints)*5 + 1
}

func main() {
	// TODO: call the function computeGrade
	fmt.Printf("Grade: %.2f\n", computeGrade(25, 30))  // 5.17
	fmt.Printf("Grade: %.2f\n", computeGrade(-80, 81)) // 5.94
	fmt.Printf("Grade: %.2f\n", computeGrade(81, 81))  // 6.00
}
