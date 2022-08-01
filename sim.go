package main

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

	canvas := NewCanvas(s.TMax * 10)

	for t := 1; t < s.TMax; t++ {
		pos := tracker.Step(target)
		canvas.DrawUpperBand(t, FloatToColor(pos.X, pos.Y, pos.Z))
		canvas.DrawLowerBand(t, FloatToColor(target.X, target.Y, target.Z))

		if t%s.NextTarget == 0 {
			ti = (ti + 1) % len(s.Targets)
			target = s.Targets[ti]
		}
	}

	canvas.WriteToFile(fileOut)
}
