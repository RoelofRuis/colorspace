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

	canvas := NewCanvas(s.TMax, 2, 10)

	for t := 1; t < s.TMax; t++ {
		pos := tracker.Step(target)
		canvas.Draw(t, 0, FloatToColor(target.X, target.Y, target.Z))
		canvas.Draw(t, 1, FloatToColor(pos.X, pos.Y, pos.Z))

		if t%s.NextTarget == 0 {
			ti = (ti + 1) % len(s.Targets)
			target = s.Targets[ti]
		}
	}

	canvas.WriteToFile("out/" + fileOut)
}
