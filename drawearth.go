package main

import (
	"fmt"
	"math"
)

func main() {
	R := 6.371e6 // radius of earth, meters
	tau := 2.0 * math.Pi
	dTheta := tau / 100.00

	for theta := 0.0; theta <= tau; theta += dTheta {
		x, y := R*math.Cos(theta), R*math.Sin(theta)
		fmt.Printf("%f\t%f\n", x, y)
	}
}
