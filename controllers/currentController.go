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

func CurrentWeather(entry *widget.Entry, tab *container.TabItem, tempUnit string, windUnit string) {
	url := "http://api.weatherapi.com/v1/current.json?key=" + os.Getenv("API_KEY") + "&q=" + entry.Text + "&aqi=yes"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error creating GET request")
	}
	if resp.StatusCode == 400 {
		tab.Content.(*fyne.Container).Objects[0].(*fyne.Container).Objects[5].(*widget.Label).SetText("THE CITY IS INVALID")
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

	vbox.Add(newCenteredText(response.Location.Name, 36))
	vbox.Add(newCenteredText(response.Location.Country, 18))
	vbox.Add(newCenteredText(strings.Split(response.Location.Localtime, " ")[1], 36))

	if response.Current.IsDay == 1 {
		vbox.Add(loadImage("weather/day/" + strings.Split(response.Current.Condition.Icon, "/")[6]))
	} else {
		vbox.Add(loadImage("weather/night/" + strings.Split(response.Current.Condition.Icon, "/")[6]))
	}

	vbox.Add(newCenteredText(response.Current.Condition.Text, 28))
	var tempText string
	if tempUnit == "°C" {
		tempText = fmt.Sprint(response.Current.TempC) + "°C"
	} else {
		tempText = fmt.Sprint(response.Current.TempF) + "°F"
	}
	vbox.Add(newCenteredText(tempText, 28))

	vbox.Add(newSpacer(0, 5))

	var windText string
	if windUnit == "Kph" {
		windText = fmt.Sprint(response.Current.WindKph) + " Kph"
	} else {
		windText = fmt.Sprint(response.Current.WindMph) + " Mph"
	}
	vb := container.NewVBox(newCenteredText("WIND", 12), loadImage("wind.svg"), newCenteredText(windText, 12))
	vb2 := container.NewVBox(newCenteredText("HUMIDITY", 12), loadImage("water_droplet.svg"), newCenteredText(strconv.Itoa(response.Current.Humidity)+" %", 12))
	vb3 := container.NewVBox(newCenteredText("PRESSURE", 12), loadImage("hpa.svg"), newCenteredText(fmt.Sprint(response.Current.PressureMb)+" hPa", 12))
	vbox.Add(newSpacer(20, 0))
	vbox.Add(container.NewCenter(container.NewHBox(vb, newSpacer(20, 0), vb2, newSpacer(20, 0), vb3)))
	vbox.Add(newSpacer(0, 5))

	pm25 := container.NewHBox(container.NewVBox(newCenteredText("PM 2.5: "+fmt.Sprint(response.Current.AirQuality.Pm25)+" µg/m³", 14), newCenteredText(airQualityPm25String(response.Current.AirQuality.Pm25), 14), newCenteredText("alarming 110 µg/m³", 12)))
	pm10 := container.NewVBox(container.NewVBox(newCenteredText("PM 10: "+fmt.Sprint(response.Current.AirQuality.Pm10)+" µg/m³", 14), newCenteredText(airQualityPm10String(response.Current.AirQuality.Pm10), 14), newCenteredText("alarming 150 µg/m³", 12)))
	vbox.Add(container.NewHBox(pm25, newSpacer(20, 0), pm10))
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

func newSpacer(w, h float32) *fyne.Container {
	r := canvas.NewRectangle(color.Transparent)
	r.SetMinSize(fyne.NewSize(w, h))
	return container.NewPadded(r)
}

func airQualityPm25String(i float64) string {
	var s strings.Builder
	s.WriteString("Air quality: ")
	if i < 13 {
		s.WriteString("Very Good")
	} else if i < 35 {
		s.WriteString("Good")
	} else if i < 55 {
		s.WriteString("Average")
	} else if i < 75 {
		s.WriteString("Sufficient")
	} else if i < 110 {
		s.WriteString("Bad")
	} else {
		s.WriteString("Very Bad")
	}
	return s.String()
}

func airQualityPm10String(i float64) string {
	var s strings.Builder
	s.WriteString("Air quality: ")
	if i < 20 {
		s.WriteString("Very Good")
	} else if i < 50 {
		s.WriteString("Good")
	} else if i < 80 {
		s.WriteString("Average")
	} else if i < 110 {
		s.WriteString("Sufficient")
	} else if i < 150 {
		s.WriteString("Bad")
	} else {
		s.WriteString("Very Bad")
	}
	return s.String()
}
