package main

import "math"

type LinearTracker struct {
	lastReading Vec3
	position    Vec3
	maxDist     float64
}

func NewLinearTracker(maxDist float64) *LinearTracker {
	return &LinearTracker{
		lastReading: Vec3{0.0, 0.0, 0.0},
		position:    Vec3{0.0, 0.0, 0.0},
		maxDist:     maxDist,
	}
}

func (t *LinearTracker) Step(target Vec3, hasMeasurement bool) []Vec3 {
	if hasMeasurement {
		t.lastReading = target
	}

	xd := t.lastReading.X - t.position.X
	yd := t.lastReading.Y - t.position.Y
	zd := t.lastReading.Z - t.position.Z

	dist := math.Sqrt(xd*xd + yd*yd + zd*zd)

	if dist < t.maxDist {
		t.position = t.lastReading
	} else {
		t.position.X += (xd * t.maxDist) / dist
		t.position.Y += (yd * t.maxDist) / dist
		t.position.Z += (zd * t.maxDist) / dist
	}

	return []Vec3{t.position}
}
