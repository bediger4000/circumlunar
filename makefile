all: spacestation_orbit.png impulsive_orbit.png nonimpulsive_orbit.png
csm: csm.go
	go build csm.go
ss: ss.go
	go build ss.go
impulse1: impulse1.go
	go build impulse1.go
drawearth: drawearth.go
	go build drawearth.go
nonimpulsive1: nonimpulsive1.go
	go build nonimpulsive1.go

earth.outline: drawearth
	./drawearth > earth.outline

spacestation_orbit.points: ss
	./ss > spacestation_orbit.points

spacestation_orbit.png: earth.outline spacestation_orbit.points orbit.load
	gnuplot < orbit.load

impulsive1.dat: impulse1
	./impulse1 > impulsive1.dat

impulsive_orbit.png: earth.outline impulsive1.dat impulse1.load
	gnuplot < impulse1.load

nonimpulsive1.dat: nonimpulsive1
	./nonimpulsive1 > nonimpulsive1.dat

nonimpulsive_orbit.png: earth.outline nonimpulsive1.dat nonimpulse1.load
	gnuplot < nonimpulse1.load
