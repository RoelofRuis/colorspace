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
	width int
}

func NewCanvas(width int) *Canvas {
	return &Canvas{
		img:   image.NewRGBA(image.Rect(0, 0, width, 120)),
		width: width,
	}
}

func (p *Canvas) DrawUpperBand(t int, col color.Color) {
	p.DrawRect(t*10, 0, (t+1)*10, 100, col)
}

func (p *Canvas) DrawLowerBand(t int, col color.Color) {
	p.DrawRect(t*10, 100, (t+1)*10, 120, col)
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
