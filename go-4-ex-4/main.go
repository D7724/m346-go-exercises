package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * (5 / 9))
}

func main() {
	fmt.Printf("Farenheit: %.2f\n", Celsius.ConvertToFahrenheit(32))
	fmt.Printf("Celcius: %.2f", Fahrenheit.ConvertToCelsius(32))
	// TODO: call the function convertCelsiusToFahrenheit
	// TODO: call the function convertFahrenheitToCelsius
}
