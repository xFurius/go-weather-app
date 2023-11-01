package model

type CurrentResponse struct {
	Location struct {
		Name      string `json:"name"`
		Region    string `json:"region"`
		Country   string `json:"country"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		IsDay     int     `json:"is_day"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		Humidity   int     `json:"humidity"`
		AirQuality struct {
			Pm25         float64 `json:"pm2_5"`
			Pm10         float64 `json:"pm10"`
			UsEpaIndex   int     `json:"us-epa-index"`
			GbDefraIndex int     `json:"gb-defra-index"`
		} `json:"air_quality"`
	} `json:"current"`
}
