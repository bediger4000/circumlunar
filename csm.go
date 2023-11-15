package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Apollo CSM:
			launch mass 63,500 lb (28,800 kg) Lunar
			dry mass 26,300 lb (11,900 kg)
			AJ10-137
			Maximum thrust	91.19 kN (20,500 lbf)
			Specific impulse	314.5 seconds (3.084 km/s)
			flow 29.5 kg/sec
			F = d(mv)/dt, F is thrust, V is exit velocity, a constant
			F = Vdm/dt
			dm/dt = F/V mass flow rate
			flow := (91190 kg m /s^2)/(3084 m/s^2) = 29.57 kg/s
	*/

	Mflow := 29.57 // propellant consumption rate, kg/sec
	dt := 0.50     // sec, time steps

	M := 28800. // start mass, kg
	V := 0.0    // initial velocity m/s
	F := 91190. // thrust kg m /s^2
	X := 0.     // starting position, m

	M0 := M
	dM := dt * Mflow

	var t float64
	fmt.Printf("# t\tMass\tX\tV\n")
	for t = 0.0; t <= 30.0; t += dt {
		fmt.Printf("%.03f\t%f\t%f\t%f\n", t, M, X, V)
		// Mave := M - dM/2.0
		// Mave := M
		// M -= dM / 2. // average mass for time period
		M -= dM
		A := F / M
		dV := dt * A
		V += dV
		dX := dt * V
		X += dX
		// M -= dM / 2.
	}
	fmt.Printf("%.03f\t%f\t%f\t%f\n", t, M, X, V)

	// deltaV := 9.82 * 314.5 * math.Log(M0/M)
	deltaV := 3084. * math.Log(M0/M)
	fmt.Printf("delta V should be %f (%f%%)\n", deltaV, 100.*deltaV/V)
}
