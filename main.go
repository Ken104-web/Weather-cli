package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type Weather struct{
    Location struct {
        Name string `json:"name"`
        Country string `json:"country"`
    }`json:"location"`
    Current struct {
        TempC float64 `json:"temp_c"`
        Condition struct{
            Text string `json:"text"`
        }`json:"condition"`
    }`json:"current"`
    Forecast struct {
        Forecastday []struct {
            Hour []struct {
                TimeEpoch int64 `json:"time_epoch"`
                TempC float64 `json:"temp_c"`
                Condition struct{
                    Text string `json:"text"`
                }`json:"condition"`
                ChanceOfRain float64 `  json:"chance_of_rain"`
                }`json:"hour"`
        }`json:"forecastday"`
    }`json:"forecast"`
}

func main(){
    res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=efd4a36bbd1e4cc6ba3145847241404&q=Kenya&days=1&aqi=no&alerts=no")
    if err != nil{
        panic(err)
    }   
    defer res.Body.Close()

    if res.StatusCode != 200 {
        panic("weather API not available")
    }

    body, err := io.ReadAll(res.Body)
    if err != nil{
        panic (err)
    }
    // fmt.Println(string(body))
    var weather Weather
    err = json.Unmarshal(body, &weather)
    if err != nil{
        panic (err)
    }
    fmt.Println(weather)

}
