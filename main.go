package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	window := myApp.NewWindow("Weather App")
	window.Resize(fyne.NewSize(500, 600))
	window.SetFixedSize(true)

	cityEntry := widget.NewEntry()
	tabs := container.NewAppTabs(
		container.NewTabItem("CURRENT WEATHER", container.NewCenter(
			container.NewVBox(
				widget.NewLabel("Enter a city:"),
				cityEntry,
				widget.NewButton("Search", nil),
			))),
		container.NewTabItem("FORECAST", widget.NewButton("test", nil)),
	)

	window.SetContent(tabs)

	window.Show()
	myApp.Run()
}
