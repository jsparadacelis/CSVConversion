package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	displacement := func(time float64) float64 {
		return 0.5*a*math.Pow(time, 2.0) + v0*time + s0
	}
	return displacement
}

func main() {

	var acc, v0, s0, time float64

	fmt.Println("Please, insert three values: ")
	fmt.Println("Acceleration value: ")
	fmt.Scan(&acc)
	fmt.Println("Initial velocity value: ")
	fmt.Scan(&v0)
	fmt.Println("Initial displacement value: ")
	fmt.Scan(&s0)

	displacement := GenDisplaceFn(acc, v0, s0)

	fmt.Println("Now, introduce time value: ")
	fmt.Scan(&time)
	fmt.Printf("\nThe displacement is: %v", displacement(time))

}
