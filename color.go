package main

type Transformer struct {
	To    func(Vec3) Vec3
	From  func(Vec3) Vec3
	Inner Tracker
}

func (t Transformer) Step(target Vec3) Vec3 {
	return t.From(t.Inner.Step(t.To(target)))
}
