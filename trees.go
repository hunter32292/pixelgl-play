package main

import (
	"image"
	"math"
	"math/rand"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

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

	// win.SetSmooth(true)

	treeSpriteSheet, err := loadPicture("trees.png")
	if err != nil {
		panic(err)
	}
	var treeFrames []pixel.Rect
	for x := treeSpriteSheet.Bounds().Min.X; x < treeSpriteSheet.Bounds().Max.X; x += 32 {
		for y := treeSpriteSheet.Bounds().Min.Y; y < treeSpriteSheet.Bounds().Max.Y; y += 32 {
			treeFrames = append(treeFrames, pixel.R(x, y, x+32, y+32))
		}
	}
	var (
		camPos       = pixel.ZV
		camSpeed     = 500.0
		camZoom      = 1.0
		camZoomSpeed = 1.2
		trees        []*pixel.Sprite
		matrices     []pixel.Matrix
	)

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			tree := pixel.NewSprite(treeSpriteSheet, treeFrames[rand.Intn(len(treeFrames))])
			trees = append(trees, tree)
			mouse := cam.Unproject(win.MousePosition())
			matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
		}
		win.Clear(colornames.Forestgreen)
		for i, tree := range trees {
			tree.Draw(win, matrices[i])
		}

		// Move left, so Camera moves X - 500
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
		}
		// Move left, so Camera moves X + 500
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
		}
		// Move left, so Camera moves y - 500
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
		}
		// Move left, so Camera moves y + 500
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
