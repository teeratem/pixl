package ui

import (
	"fyne.io/fyne/v2"
	"github.com/teeratem/pixl/apptype"
	"github.com/teeratem/pixl/pxcanvas"
	"github.com/teeratem/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
