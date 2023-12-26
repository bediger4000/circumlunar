package main

/*
 * Simulate 1075 mi circular orbit for 230 days to ensure that the
 * numeric integration doesn't lose too much energy.
 */

import (
	"fmt"
	"math"
)

func main() {

	const G = 6.673e-11 // m^3/(kg s^2)

	Mearth := 5.97e24

	// Initial conditions - 1075 mi circular orbit
	X := 6.371e6 + 1730044.745 // meters
	Y := 0.0
	Vx := 0.0
	Vy := 8500. // meters/second
	GM1 := G * Mearth

	// 8.101e6 m orbit radius
	// orbit circumference = 2*pi*8.101e6 = 5.09E7
	// t = 5.09e7/7012.5 = 7258 sec
	var t, r float64
	dt := .250 // seconds

	// vis viva eqn: v^2 = GM(2/r - 1/a), a is semi-major axis
	// v^2 = GM/a[(2a-r)/r]
	// for V = 7069.5 m/s, at radius of 8101044.6 m,
	// a = 8101044.745 m
	a := 8101044.745
	a2 := 2.0 * a
	GMa := GM1 / a

	loopCounter := 0

	fmt.Printf("# t\tVx\tVy\tx\ty\tr\n")
	for t = 0.0; t <= 20000000.0; t += dt {
		loopCounter++

		r2 := X*X + Y*Y
		r = math.Sqrt(r2)

		vv := GMa * ((a2 - r) / r)

		if (loopCounter % 200) == 1 {
			fmt.Printf("%f\t%f\t%f\t%f\t%f\t%f\t%f\n", t, Vx, Vy, X, Y, r, (Vx*Vx+Vy*Vy)/vv)
		}

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
