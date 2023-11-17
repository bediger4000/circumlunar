all: csm ss earthoutline
csm: csm.go
	go build csm.go
ss: ss.go
	go build ss.go
drawearth: drawearth.go
	go build drawearth.go

earth.ouline: drawearth
	./drawearth > earth.outline

spacestation_orbit.points: ss
	./ss > spacestation_orbit.points

spacestation_orbit.png: earth.ouline spacestation_orbit.points orbit.load
	gnuplot < orbit.load
