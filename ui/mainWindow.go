package ui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const user string = "Max"

func Create(w fyne.Window) {
	// window should always be full screen!
	w.SetFullScreen(true)

	w.SetContent(container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), createGreeting(), layout.NewSpacer()),
		layout.NewSpacer(),
		createWeatherLayout(),
		layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), createCurrentDayLayout(), layout.NewSpacer()),
		layout.NewSpacer(),
		createTodoList(),
		layout.NewSpacer(),
	))
}

func createGreeting() *canvas.Text {
	return canvas.NewText("Good day, "+user+"!", color.White)
}

func createWeatherLayout() *fyne.Container {
	var weatherImage *canvas.Image = canvas.NewImageFromResource(fyne.NewStaticResource("sunny.png", ResourceSunnyPng.StaticContent))
	weatherImage.FillMode = canvas.ImageFillOriginal
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(),
		weatherImage,
		layout.NewSpacer(),
		container.NewGridWithRows(2, canvas.NewText("22 °C", color.White), widget.NewLabel("5 °C")),
		layout.NewSpacer(),
	)
}

func createCurrentDayLayout() *canvas.Text {
	return canvas.NewText(time.Now().Format(time.RFC1123), color.White)
}

func createTodoList() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		canvas.NewLine(color.White),
		layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), canvas.NewText("TODO", color.White), layout.NewSpacer()),
		layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewLabel("Wash dishes"), layout.NewSpacer()),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewLabel("take umbrella"), layout.NewSpacer()),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewLabel("Go for a walk!"), layout.NewSpacer()),
		layout.NewSpacer(),
	)
}
