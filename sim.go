package main

import (
	"math/rand"
	"time"
)

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

	canvas := NewCanvas(s.TMax, tracker.NumLeds()+1, 10)

	rand.Seed(0)
	jitter := 50.0

	for t := 1; t < s.TMax; t++ {
		target.X += (rand.Float64() - 0.5) * jitter
		target.Y += (rand.Float64() - 0.5) * jitter
		target.Z += (rand.Float64() - 0.5) * jitter

		canvas.Draw(t, 0, FloatToColor(target.X, target.Y, target.Z))

		for i, pos := range tracker.Step(target) {
			canvas.Draw(t, i+1, FloatToColor(pos.X, pos.Y, pos.Z))
		}

		if t%s.NextTarget == 0 {
			ti = (ti + 1) % len(s.Targets)
			target = s.Targets[ti]
		}
	}

	canvas.WriteToFile("out/" + fileOut + time.Now().Format("20060102150405") + ".png")
}
