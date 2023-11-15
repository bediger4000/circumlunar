package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	 * Wernher von Braun's 1075 mile polar orbit
	 * r =
	 * Vtan =
	 */

	const G = 6.673e-11 // m^3/(kg s^2)

	Mearth := 5.97e24
	// Msatellite := 1.0

	X := 6.371e6 + 1730044.745 // meters
	Y := 0.0
	Vx := 0.0
	// 15840 mi/hr * 5280 ft/mi * 1/3.28 m/ft * 1/3600 hr/sec = 7083 m/s
	// Vy := 7082.9 // meters/second
	Vy := 7069.4937 // meters/second
	GM1 := G * Mearth

	var t, r float64
	dt := .250 // seconds
	fmt.Printf("# t\tVx\tVy\tx\ty\tr\n")
	for t = 0.0; t <= 7485.; t += dt {

		r2 := X*X + Y*Y
		r = math.Sqrt(r2)
		fmt.Printf("%f\t%f\t%f\t%f\t%f\t%f\t%f\n", t, Vx, Vy, X, Y, r, r/8101487.804800)

		// magnitude of attraction F = G*M1*m2/(r^2)
		Fmag := GM1 / r2
		Fx := (-X / r) * Fmag
		Fy := (-Y / r) * Fmag

		Ax := Fx
		Ay := Fy

		dVx := Ax * dt
		dVy := Ay * dt

		Vx += dVx
		Vy += dVy

		X += Vx * dt
		Y += Vy * dt

	}
}
