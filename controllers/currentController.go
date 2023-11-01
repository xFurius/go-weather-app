package controllers

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"weather-app/main/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CurrentWeather(entry *widget.Entry, tab *container.TabItem) {
	url := "http://api.weatherapi.com/v1/current.json?key=" + os.Getenv("API_KEY") + "&q=" + entry.Text + "&aqi=yes"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error creating GET request")
	}
	if resp.StatusCode == 400 { //
		tab.Content.(*fyne.Container).Objects[0].(*fyne.Container).Objects[3].(*widget.Label).SetText("THE CITY IS INVALID")
		runtime.Goexit()
		return
	}

	var response *model.CurrentResponse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read resp body")
	}
	resp.Body.Close()

	if err = json.Unmarshal(respBody, &response); err != nil {
		log.Fatal("Failed to unmarshal", err)
	}

	tab.Content.(*fyne.Container).RemoveAll()
	tab.Content.(*fyne.Container).Add(container.NewVBox())
	vbox := tab.Content.(*fyne.Container).Objects[0].(*fyne.Container)

	//TODO: automation?
	vbox.Add(newCenteredText(response.Location.Name, 36))
	vbox.Add(newCenteredText(response.Location.Country, 18))
	vbox.Add(newCenteredText(strings.Split(response.Location.Localtime, " ")[1], 36))

	//TODO: change resource name based on the weather
	vbox.Add(loadImage("113.png"))

	vbox.Add(newCenteredText(response.Current.Condition.Text, 28))
	//TODO: give user an option to choose C or F
	vbox.Add(newCenteredText(fmt.Sprint(response.Current.TempC)+"°C / "+fmt.Sprint(response.Current.TempF)+"°F", 28))

	r := canvas.NewRectangle(color.Transparent)
	r.SetMinSize(fyne.NewSize(0, 20))
	vbox.Add(container.NewPadded(r))

	//TODO: make it better
	//TODO: give user an option to choose kph or mph
	vb := container.NewVBox(newCenteredText("WIND", 12), loadImage("wind.svg"), newCenteredText(fmt.Sprint(response.Current.WindKph)+" kph", 12))
	vb2 := container.NewVBox(newCenteredText("HUMIDITY", 12), loadImage("water_droplet.svg"), newCenteredText(strconv.Itoa(response.Current.Humidity)+" %", 12))
	vb3 := container.NewVBox(newCenteredText("PRESSURE", 12), loadImage("hpa.svg"), newCenteredText(fmt.Sprint(response.Current.PressureMb)+" hPa", 12))
	//TODO: add spacing between vbs
	r = canvas.NewRectangle(color.Transparent)
	r.SetMinSize(fyne.NewSize(20, 0))
	vbox.Add(container.NewPadded(r))
	vbox.Add(container.NewHBox(vb, r, vb2, r, vb3))

	tab.Content.Refresh()
}

func newCenteredText(c string, s float32) *canvas.Text {
	a := canvas.NewText(c, color.White)
	a.TextSize = s
	a.Alignment = fyne.TextAlignCenter
	return a
}

func loadImage(name string) *canvas.Image {
	res, err := fyne.LoadResourceFromPath("./resources/" + name)
	if err != nil {
		log.Println(err)
	}
	image := canvas.NewImageFromResource(res)
	image.FillMode = canvas.ImageFillOriginal
	return image
}
