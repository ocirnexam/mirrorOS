package main

import (
	"mirrorApp/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&(ui.CustomTheme{}))
	w := a.NewWindow("Mirror App")
	ui.Create(w)
	w.ShowAndRun()
}
