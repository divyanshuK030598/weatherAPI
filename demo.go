package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WeatherDetails struct {
	Data         []WeatherData `json:"data"`
	City_name    string        `json:"city_name"`
	Lon          float32       `json:"lon"`
	Timezone     string        `json:"timezone"`
	Lat          float32       `json:"lat"`
	Country_code string        `json:"country_code"`
	State_code   string        `json:"state_code"`
}
type WeatherData struct {
	Timestamp_utc   string  `json:"timestamp_utc"`
	Snow            float64 `json:"snow"`
	Temp            float32 `json:"temp"`
	Timestamp_local string  `json:"timestamp_local"`
	Ts              int64   `json:"ts"`
	Precip          float32 `json:"precip"`
}

func main() {
	fmt.Println("Enter Your Latitute: ")
	var latitude string
	fmt.Scanln(&latitude)
	fmt.Println("Enter Your longitude: ")
	var longitude float32
	fmt.Scanln(&longitude)
	url := fmt.Sprintf("https://weatherbit-v1-mashape.p.rapidapi.com/forecast/minutely?lat=%v&lon=%v", latitude, longitude)
	//url := "https://weatherbit-v1-mashape.p.rapidapi.com/forecast/minutely?lat=35.5&lon=-78.5"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "539c1c43f8msh1b147a3761cf2eap178220jsnba861e9188de")
	req.Header.Add("x-rapidapi-host", "weatherbit-v1-mashape.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject WeatherDetails
	json.Unmarshal(body, &responseObject)
	fmt.Println("City_name", responseObject.City_name)
	fmt.Println("Latitude", responseObject.Lat)
	fmt.Println("Longitude", responseObject.Lon)
	fmt.Println("Timezone", responseObject.Timezone)
	fmt.Println("State_Code", responseObject.State_code)
	fmt.Println("Country_Code", responseObject.Country_code)
	fmt.Println("Weather data at different time of day")
	for _, v := range responseObject.Data {
		fmt.Println("Timespamt_utc : ", v.Timestamp_utc)
		fmt.Println("Snow : ", v.Snow)
		fmt.Println("Temperature : ", v.Temp)
		fmt.Println("Timestamp_local : ", v.Timestamp_local)
		fmt.Println("Ts : ", v.Ts)
		fmt.Println("Precipitation : ", v.Precip)
		fmt.Println()

	}

}
