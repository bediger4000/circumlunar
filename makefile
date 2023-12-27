all: spacestation_orbit.png impulsive_orbit.png nonimpulsive_orbit.png \
	orbits_overlay.png circularize.png hohmann_overlay.png

clean:
	-rm -rf spacestation_orbit.png impulsive_orbit.png \
		nonimpulsive_orbit.png orbits_overlay.png circularize.png \
		hohmann_overlay.png \
		earth.outline moon.outline \
		impulsive1.dat impulsive2.dat circularize.dat \
		nonimpulsive1.dat nonimpulsive2.dat \
		pastmoon.dat spacestation_orbit.dat \
		csn ss impulse1 impusle2 drawearth drawmoon \
		nonimpulsive1 nonimpulsive2 \
		pastmoon ss impulse2 circularize


csm: csm.go
	go build csm.go
ss: ss.go
	go build ss.go
impulse1: impulse1.go
	go build impulse1.go
impulse2: impulse2.go
	go build impulse2.go
drawearth: drawearth.go
	go build drawearth.go
drawmoon: drawmoon.go
	go build drawmoon.go
nonimpulsive1: nonimpulsive1.go
	go build nonimpulsive1.go
nonimpulsive2: nonimpulsive2.go
	go build nonimpulsive2.go
pastmoon: pastmoon.go
	go build pastmoon.go
circularize: circularize.go
	go build circularize.go

earth.outline: drawearth
	./drawearth > earth.outline
moon.outline: drawmoon
	./drawmoon > moon.outline

spacestation_orbit.dat: ss
	./ss > spacestation_orbit.dat

spacestation_orbit.png: earth.outline spacestation_orbit.dat orbit.load
	gnuplot < orbit.load

impulsive1.dat: impulse1
	./impulse1 > impulsive1.dat

impulsive2.dat: impulse2
	./impulse2 > impulsive2.dat

pastmoon.dat: pastmoon
	./pastmoon > pastmoon.dat

circularize.dat: circularize
	./circularize > circularize.dat


impulsive_orbit.png: earth.outline impulsive1.dat impulse1.load
	gnuplot < impulse1.load

nonimpulsive1.dat: nonimpulsive1
	./nonimpulsive1 > nonimpulsive1.dat

nonimpulsive2.dat: nonimpulsive2
	./nonimpulsive2 > nonimpulsive2.dat

nonimpulsive_orbit.png: earth.outline nonimpulsive1.dat nonimpulse1.load
	gnuplot < nonimpulse1.load

orbits_overlay.png: earth.outline nonimpulsive1.dat impulsive1.dat overlay.load
	gnuplot < overlay.load

circularize.png: circularize.dat circularize.load moon.outline earth.outline
	gnuplot < circularize.load

hohmann_overlay.png: overlay2.load earth.outline impulsive2.dat nonimpulsive2.dat
	gnuplot < overlay2.load
