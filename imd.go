package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Triangle
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(200, 100)) // Bottom Left
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(800, 100)) // Bottom Right
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(500, 700)) // Top Center
	imd.Polygon(0)

	// Drawing lines
	// imd := imdraw.New(nil)
	// imd.Color = colornames.Blueviolet
	// imd.EndShape = imdraw.RoundEndShape
	// imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	// imd.EndShape = imdraw.SharpEndShape
	// imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	// imd.Line(30)

	// Drawing Circles, Arch, Ellipse
	// imd.Color = colornames.Limegreen
	// imd.Push(pixel.V(500, 500))
	// imd.Circle(300, 50)
	// imd.Color = colornames.Navy
	// imd.Push(pixel.V(200, 500), pixel.V(800, 500))
	// imd.Ellipse(pixel.V(120, 80), 0)
	// imd.Color = colornames.Red
	// imd.EndShape = imdraw.RoundEndShape
	// imd.Push(pixel.V(500, 350))
	// imd.CircleArc(150, -math.Pi, 0, 30)

	// circle := imdraw.New(nil)
	// circle.Color = colornames.Red
	// circle.Push(pixel.V(100, 100))
	// circle.Circle(64, 0)

	for !win.Closed() {

		win.Clear(colornames.White)
		// circle.Draw(win)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
