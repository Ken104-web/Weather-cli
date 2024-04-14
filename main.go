package main

import "net/http"



func main(){
    res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=efd4a36bbd1e4cc6ba3145847241404&q=London&aqi=no")
    if err != nil{
        panic(err)
    }   
    defer res.Body.Close()

    if res.StatusCode != 200 {
        panic("weather API not available")
    }
}
