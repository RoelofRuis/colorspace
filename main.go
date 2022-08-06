package main

import "github.com/hisamafahri/coco"

type Tracker interface {
	Step(target Vec3) Vec3
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
		})

	forceTracker := NewForceTracker(8.0, 0.99)
	linTracker := NewLinearTracker(10.0)

	colTransformer := Transformer{
		To: func(in Vec3) Vec3 {
			out := coco.Rgb2Xyz(in.X, in.Y, in.Z)
			return Vec3{X: out[0], Y: out[1], Z: out[2]}
		},
		From: func(in Vec3) Vec3 {
			out := coco.Xyz2Rgb(in.X, in.Y, in.Z)
			return Vec3{X: out[0], Y: out[1], Z: out[2]}
		},
		Inner: NewLinearTracker(3.0),
	}

	sim.Run(colTransformer, "col_transformer.png")
	sim.Run(forceTracker, "force_out.png")
	sim.Run(linTracker, "lin_out.png")
}
