package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const user string = "Max"

func Create(w fyne.Window) {
	// window should always be full screen!
	w.SetFullScreen(true)

	// create a greeting label
	var heading *widget.Label = widget.NewLabel("Good day, " + user + "!")
	var centerBox *fyne.Container = container.NewCenter()
	var verticalBox *fyne.Container = container.NewVBox()
	verticalBox.Add(heading)
	verticalBox.Add(widget.NewButton("Show Mirror!", func() {
		heading.SetText("ABC")
	}))
	verticalBox.Add(createWeatherLayout())
	centerBox.Add(verticalBox)
	w.SetContent(centerBox)
}

func createWeatherLayout() *fyne.Container {
	return container.NewGridWithColumns(4, widget.NewLabel("Temperature"), widget.NewLabel("Wind"), widget.NewLabel("Weather"), widget.NewLabel("Time"))
}
