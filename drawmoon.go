package main

// Draw the moon to scale, out where it should be,
// if earth's center is at (0,0)

import (
	"fmt"
	"math"
)

func main() {
	R := 1737.4e3 // radius of moon, meters
	tau := 2.0 * math.Pi
	dTheta := tau / 100.00

	for theta := 0.0; theta <= tau; theta += dTheta {
		x, y := R*math.Cos(theta), R*math.Sin(theta)
		fmt.Printf("%f\t%f\n", x-3.844e8, y)
	}
}
