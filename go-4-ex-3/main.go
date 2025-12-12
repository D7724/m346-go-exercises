package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}
func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	d := computeDiscriminant(a, b, c)
	var values = make([]float64, 0, 2)
	if d > 0 {
		values = append(values, (-b+d)/(2*a))
		values = append(values, (-b-d)/(2*a))
	} else if d == 0 {
		values = append(values, (-b+d)/(2*a))
	}
	return values
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Printf("D+ %.2f\n", computeQuadraticFormula(3, 4, 1))
	fmt.Printf("D0 %.2f\n", computeQuadraticFormula(2, 4, 2))
	fmt.Printf("D- %.2f\n", computeQuadraticFormula(3, 4, 2))
}
