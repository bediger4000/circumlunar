package main

/*
 * Wernher von Braun's 1075 mile polar orbit
 */

import (
	"fmt"
	"math"
)

func main() {
	/*
	 * Wernher von Braun's 1075 mile polar orbit
	 * r = 8101044.745 m
	 * Vtan = sqrt(GM/r) = sqrt(6.673e-11*5.97e24/8101044.745) = 7012.6 m/s
	 * 8.101e6 m orbit radius
	 * orbit circumference = 2*pi*8.101e6 = 5.09E7
	 * time to complete a single orbig 5.09e7/7012.5 = 7258 sec
	 */

	const G = 6.673e-11 // m^3/(kg s^2)

	Mearth := 5.97e24

	X := 6.371e6 + 1730044.745 // meters
	Y := 0.0
	Vx := 0.0
	Vy := 7012.6 // meters/second, this is a bit different than von Braun's figure
	GM1 := G * Mearth

	a := 8101044.745
	a2 := 2.0 * a
	GMa := GM1 / a

	// distance traveled
	S := 0.0

	var t, r float64
	dt := .250 // seconds
	fmt.Printf("# t\tVx\tVy\tx\ty\tv*v/vv\tr\tr/radius\tS\n")
	for t = 0.0; t <= 7300.; t += dt {

		r2 := X*X + Y*Y
		r = math.Sqrt(r2)
		//  v^2 according to vis-viva eqn
		vv := GMa * ((a2 - r) / r)
		fmt.Printf("%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\t%f\n",
			t, Vx, Vy, X, Y, (Vx*Vx+Vy*Vy)/vv, r, r/8101044.745, S)

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

		// Symplectic method - https://www.mgaillard.fr/2021/07/11/euler-integration.html
		dX := Vx * dt
		dY := Vy * dt

		// keep track of distance traveled
		dS := math.Sqrt(dX*dX + dY*dY)
		S += dS

		X += dX
		Y += dY

	}
}
