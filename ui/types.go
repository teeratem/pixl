package ui

import (
	"fyne.io/fyne/v2"
	"github.com/teeratem/pixel/apptype"
	"github.com/teeratem/pixel/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
