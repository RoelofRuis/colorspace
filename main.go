package main

import "math/rand"

type Tracker interface {
	NumLeds() int
	Step(target Vec3) []Vec3
}

func main() {
	rand.Seed(0)

	sim := NewSimulation(
		25,
		10.0,
		[]Vec3{
			{0.0, 0.0, 250.0},
			{0.0, 0.0, 0.0},
			{250.0, 250.0, 0.0},
			{0.0, 0.0, 0.0},
			{250.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 250.0, 0.0},
			{0.0, 0.0, 0.0},
			{250.0, 0.0, 250.0},
			{0.0, 0.0, 0.0},
			{0.0, 250.0, 250.0},
			{0.0, 0.0, 0.0},
		})

	linTracker := NewLinearTracker(
		20,
		10.0,
		20,
	)
	sim.Run(linTracker, "lin_out")
}
