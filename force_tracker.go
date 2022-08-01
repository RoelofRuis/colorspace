package main

import "math"

type ForceTracker struct {
	lastReading  Vec3
	position     Vec3
	speed        Vec3
	acceleration Vec3
	alpha        float64
	damping      float64
}

func NewForceTracker(alpha, damping float64) *ForceTracker {
	return &ForceTracker{
		lastReading:  Vec3{0.0, 0.0, 0.0},
		position:     Vec3{0.0, 0.0, 0.0},
		speed:        Vec3{0, 0, 0},
		acceleration: Vec3{0, 0, 0},
		alpha:        alpha,
		damping:      damping,
	}
}

func (t *ForceTracker) Step(target Vec3) Vec3 {
	t.position.X += t.speed.X
	t.position.Y += t.speed.Y
	t.position.Z += t.speed.Z

	t.speed.X += t.acceleration.X
	t.speed.Y += t.acceleration.Y
	t.speed.Z += t.acceleration.Z

	xd := 0.0
	yd := 0.0
	zd := 0.0
	l := 0.0
	if !target.IsZero() {
		t.lastReading = target
	}
	xd = t.lastReading.X - t.position.X
	yd = t.lastReading.Y - t.position.Y
	zd = t.lastReading.Z - t.position.Z
	l = 1 / math.Sqrt(xd*xd+yd*yd+zd*zd)

	t.acceleration.X = xd*l*t.alpha - t.speed.X*t.damping
	t.acceleration.Y = yd*l*t.alpha - t.speed.Y*t.damping
	t.acceleration.Z = zd*l*t.alpha - t.speed.Z*t.damping

	return t.position
}
