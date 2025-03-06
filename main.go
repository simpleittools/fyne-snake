package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func setupGame() *fyne.Container {
	var segments []fyne.CanvasObject

	for i := 0; i < 10; i++ {
		r := canvas.NewRectangle(color.NRGBA{0x00, 0xff, 0, 0xff})
		r.Resize(fyne.NewSize(10, 10))
		r.Move(fyne.NewPos(90, float32(50+i*10)))
		segments = append(segments, r)
	}
	return container.NewWithoutLayout(segments...)
}

func main() {
	a := app.New()

	w := a.NewWindow("Snake")
	w.SetContent(setupGame())
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
