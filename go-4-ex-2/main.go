package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(shortSides ShortSides) float64 {
	return math.Sqrt(math.Pow(2, shortSides.a) + math.Pow(2, shortSides.b))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Printf("%.2f\n", computeHypotenuse(ShortSides{2, 4}))
}
