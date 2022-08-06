package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

type Canvas struct {
	img   *image.RGBA
	scale int
}

func NewCanvas(width, height, scale int) *Canvas {
	return &Canvas{
		img:   image.NewRGBA(image.Rect(0, 0, width*scale, height*scale)),
		scale: scale,
	}
}

func (p *Canvas) Draw(x, y int, col color.Color) {
	p.DrawRect(x*p.scale, y*p.scale, (x+1)*p.scale, (y+1)*p.scale, col)
}

func (p *Canvas) DrawRect(x0, y0, x1, y1 int, col color.Color) {
	draw.Draw(
		p.img,
		image.Rect(x0, y0, x1, y1),
		image.NewUniform(col),
		image.Point{},
		draw.Src,
	)
}

func (p *Canvas) WriteToFile(fname string) {
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, p.img); err != nil {
		panic(err)
	}
}

func FloatToColor(r, g, b float64) color.Color {
	return color.RGBA{
		ClippedUint8(r),
		ClippedUint8(g),
		ClippedUint8(b),
		255,
	}
}

func ClippedUint8(f float64) uint8 {
	r := math.Round(f)
	if r < 0 {
		return uint8(0)
	}
	if r > 255.0 {
		return uint8(255)
	}
	return uint8(r)
}
