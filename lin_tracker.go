package main

import "math"

type LinearTracker struct {
	lastReading Vec3
	averagesIdx int
	averages    []Vec3
	leds        []Vec3
	maxDist     float64
}

func NewLinearTracker(numLeds int, maxDist float64, numAverages int) *LinearTracker {
	var leds = make([]Vec3, numLeds)
	for i := 0; i < numLeds; i++ {
		leds[i] = Vec3{0.0, 0.0, 0.0}
	}

	return &LinearTracker{
		lastReading: Vec3{0.0, 0.0, 0.0},
		averagesIdx: 0,
		averages:    make([]Vec3, numAverages),
		leds:        leds,
		maxDist:     maxDist,
	}
}

func (t *LinearTracker) Step(target Vec3) []Vec3 {
	hasReading := target.X+target.Y+target.Z > 200

	if hasReading {
		t.averages[t.averagesIdx] = target
		t.averagesIdx = (t.averagesIdx + 1) % len(t.averages)

		avg := Vec3{}
		for _, average := range t.averages {
			avg.X += average.X
			avg.Y += average.Y
			avg.Z += average.Z
		}

		t.lastReading = Vec3{
			avg.X / float64(len(t.averages)),
			avg.Y / float64(len(t.averages)),
			avg.Z / float64(len(t.averages)),
		}
	}

	xd := t.lastReading.X - t.leds[0].X
	yd := t.lastReading.Y - t.leds[0].Y
	zd := t.lastReading.Z - t.leds[0].Z
	dist := math.Sqrt(xd*xd + yd*yd + zd*zd)

	var next Vec3 = t.leds[0]
	if dist < t.maxDist {
		next = t.lastReading
	} else {
		next.X += (xd * t.maxDist) / dist
		next.Y += (yd * t.maxDist) / dist
		next.Z += (zd * t.maxDist) / dist
	}
	t.leds = append([]Vec3{next}, t.leds[0:len(t.leds)-1]...)

	return t.leds
}

func (t *LinearTracker) NumLeds() int {
	return len(t.leds)
}
