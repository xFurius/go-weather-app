package main

import (
	"os"
	"weather-app/main/controllers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

var currentTab *container.TabItem

func init() {
	godotenv.Load(".env")
	os.Setenv("FYNE_FONT", "resources/Virgil GS Regular.ttf")
}

func main() {
	myApp := app.New()

	window := myApp.NewWindow("Weather App")
	window.Resize(fyne.NewSize(500, 600))
	window.SetFixedSize(true)

	cityEntry := widget.NewEntry()
	tempUnit := widget.NewRadioGroup([]string{"°C", "°F"}, nil)
	tempUnit.Horizontal = true
	windUnit := widget.NewRadioGroup([]string{"Kph", "Mph"}, nil)
	windUnit.Horizontal = true
	currentTab = container.NewTabItem("CURRENT WEATHER", container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Enter a city:"),
			cityEntry,
			tempUnit,
			windUnit,
			widget.NewButton("Search", func() {
				go controllers.CurrentWeather(cityEntry, currentTab, tempUnit.Selected, windUnit.Selected)
			}),
			widget.NewLabel(""),
		)))

	tabs := container.NewAppTabs(
		currentTab,
		container.NewTabItem("FORECAST", widget.NewButton("test", nil)),
	)

	window.SetContent(tabs)

	window.Show()
	myApp.Run()
}
