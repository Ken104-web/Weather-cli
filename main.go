package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `  json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
    q := "Kenya"
    if len(os.Args) >= 2 {
        q = os.Args[1]
    }
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=efd4a36bbd1e4cc6ba3145847241404&q=" + q + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	// fmt.Println(weather)
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s: %.0fC, %s",
        location.Name,
        location.Country,
        current.TempC,
        current.Condition.Text,
    )
    for _, hour := range hours {
        date := time.Unix(hour.TimeEpoch, 0)
        if date.Before(time.Now()) {
            continue
        }
        fmt.Printf("%s, %.0fC%%, %0f, %s\n",
            date.Format("15.04"),
            hour.TempC,
            hour.ChanceOfRain,
            hour.Condition.Text,
        )
    }
}
