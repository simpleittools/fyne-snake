package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

// snakePart is the game state
type snakePart struct {
	x, y float32
}

var (
	snakeParts []snakePart
	game       *fyne.Container
)

func setupGame() *fyne.Container {
	var segments []fyne.CanvasObject

	for i := 0; i < 10; i++ {
		r := canvas.NewRectangle(color.NRGBA{0x00, 0xff, 0, 0xff})
		r.Resize(fyne.NewSize(10, 10))
		r.Move(fyne.NewPos(90, float32(50+i*10)))
		segments = append(segments, r)

		seg := snakePart{
			9, float32(5 + i),
		}
		snakeParts = append(snakeParts, seg)
	}
	return container.NewWithoutLayout(segments...)
}

func refreshGame() {
	for i, seg := range snakeParts {
		rect := game.Objects[i]
		rect.Move(fyne.NewPos(seg.x*10, seg.y*10))
	}
	game.Refresh()
}

func runGame() {
	for {
		time.Sleep(time.Millisecond * 250)
		for i := len(snakeParts) - 1; i >= 1; i-- {
			snakeParts[i] = snakeParts[i-1]
		}
		snakeParts[0].y--
		refreshGame()
	}
}

func main() {
	a := app.New()

	w := a.NewWindow("Snake")
	game = setupGame()
	w.SetContent(game)
	go runGame()
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
