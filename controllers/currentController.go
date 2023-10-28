package controllers

import (
	"encoding/json"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"weather-app/main/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CurrentWeather(entry *widget.Entry, tab *container.TabItem) {
	log.Println(entry.Text, os.Getenv("API_KEY"))

	url := "http://api.weatherapi.com/v1/current.json?key=" + os.Getenv("API_KEY") + "&q=" + entry.Text + "&aqi=yes"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error creating GET request")
	}
	log.Println(resp.StatusCode)

	var response *model.CurrentResponse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read resp body")
	}
	resp.Body.Close()

	if err = json.Unmarshal(respBody, &response); err != nil {
		log.Fatal("Failed to unmarshal", err)
	}

	log.Println(response.Location.Name)

	// currentTab = container.NewTabItem("CURRENT WEATHER", container.NewCenter(
	// 	container.NewVBox(
	// 		widget.NewLabel("Enter a city:"), <---
	// 		cityEntry,
	// 		widget.NewButton("Search", func() {
	// 			go controllers.CurrentWeather(cityEntry, currentTab)
	// 		}),
	// 	)))

	tab.Content.(*fyne.Container).RemoveAll()
	tab.Content.(*fyne.Container).Add(container.NewVBox())
	vbox := tab.Content.(*fyne.Container).Objects[0].(*fyne.Container)

	a := canvas.NewText(response.Location.Name, color.White)
	a.TextSize = 32
	a.Alignment = fyne.TextAlignCenter
	vbox.Add(a)

	a = canvas.NewText(response.Location.Country, color.White)
	a.TextSize = 22
	a.Alignment = fyne.TextAlignCenter
	vbox.Add(a)

	vbox.Add(widget.NewLabel("test3"))
	tab.Content.Refresh()
}
