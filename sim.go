package main

import "math/rand"

type Simulation struct {
	TMax       int
	NextTarget int
	Targets    []Vec3
}

func NewSimulation(nextTarget int, targets []Vec3) *Simulation {
	return &Simulation{
		TMax:       nextTarget * len(targets),
		NextTarget: nextTarget,
		Targets:    targets,
	}
}

func (s Simulation) Run(tracker Tracker, fileOut string) {
	ti := 0
	target := s.Targets[ti]
	hasMeasurement := false

	canvas := NewCanvas(s.TMax, 2, 50)

	rand.Seed(0)
	jitter := 20.0

	for t := 1; t < s.TMax; t++ {
		target.X += (rand.Float64() - 0.5) * jitter
		target.Y += (rand.Float64() - 0.5) * jitter
		target.Z += (rand.Float64() - 0.5) * jitter

		hasMeasurement = target.X+target.Y+target.Z > 50

		canvas.Draw(t, 0, FloatToColor(target.X, target.Y, target.Z))

		for i, pos := range tracker.Step(target, hasMeasurement) {
			canvas.Draw(t, i+1, FloatToColor(pos.X, pos.Y, pos.Z))
		}

		if t%s.NextTarget == 0 {
			ti = (ti + 1) % len(s.Targets)
			target = s.Targets[ti]
		}
	}

	canvas.WriteToFile("out/" + fileOut)
}
