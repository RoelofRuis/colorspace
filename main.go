package main

type Tracker interface {
	Step(target Vec3, hasMeasurement bool) []Vec3
}

func main() {
	sim := NewSimulation(
		25,
		[]Vec3{
			{0.0, 0.0, 250.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{250.0, 250.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{250.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 250.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{250.0, 0.0, 250.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 250.0, 250.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
		})

	linTracker := NewLinearTracker(10.0)
	sim.Run(linTracker, "lin_out.png")
}
