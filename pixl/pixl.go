package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/teeratem/pixel/apptype"
	"github.com/teeratem/pixel/swatch"
	"github.com/teeratem/pixel/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()

}
