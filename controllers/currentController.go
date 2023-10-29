package controllers

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	//TODO: error handling if city is invalid
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

	tab.Content.(*fyne.Container).RemoveAll()
	tab.Content.(*fyne.Container).Add(container.NewVBox())
	vbox := tab.Content.(*fyne.Container).Objects[0].(*fyne.Container)

	//TODO: automation?
	vbox.Add(newText(response.Location.Name, 36))
	vbox.Add(newText(response.Location.Country, 18))
	vbox.Add(newText(strings.Split(response.Location.Localtime, " ")[1], 36))

	//TODO: change resource name based on the weather
	res, err := fyne.LoadResourceFromPath("./resources/113.png")
	log.Println(err)

	image := canvas.NewImageFromResource(res)
	image.FillMode = canvas.ImageFillOriginal
	vbox.Add(image)

	vbox.Add(newText(response.Current.Condition.Text, 28))
	//TODO: give user an option to choose C or F
	vbox.Add(newText(fmt.Sprint(response.Current.TempC)+"°C / "+fmt.Sprint(response.Current.TempF)+"°F", 28))

	//TODO: make it better
	vb := container.NewVBox(widget.NewLabel("NE"), widget.NewLabel("NE"), widget.NewLabel("NE"))
	vb2 := container.NewVBox(widget.NewLabel("NE"), widget.NewLabel("NE"), widget.NewLabel("NE"))
	vb3 := container.NewVBox(widget.NewLabel("NE"), widget.NewLabel("NE"), widget.NewLabel("NE"))
	vbox.Add(container.NewHBox(vb, vb2, vb3))

	tab.Content.Refresh()
}

func newText(c string, s float32) *canvas.Text {
	a := canvas.NewText(c, color.White)
	a.TextSize = s
	a.Alignment = fyne.TextAlignCenter
	return a
}
