package main

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) IsZero() bool {
	return v.X == 0.0 && v.Y == 0.0 && v.Z == 0.0
}
