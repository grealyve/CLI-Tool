/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type WeatherForecast struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity int16 `json: "humidity"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json: "hour"`
		} `json: "forecastday"`
	} `json: "forecast"`
}

// weatherCmd represents the weather command
var WeatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		q := "Mugla"

		if len(args) >= 0 {
			q = args[0]
		}

		res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=f5270bfce5e247a398163148232007&q=" + q + "&days=3&aqi=no&alerts=no")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			panic("Weather API not available")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var weatherF WeatherForecast
		err = json.Unmarshal(body, &weatherF)
		if err != nil {
			panic(err)
		}

		location, current, hours := weatherF.Location, weatherF.Current, weatherF.Forecast.Forecastday[0].Hour

		fmt.Printf("%s, %s: %.0fC, %s, Humidity: %%%v\n", location.Name, location.Country, current.TempC, current.Condition.Text, current.Humidity)

		for _, hour := range hours {
			date := time.Unix(hour.TimeEpoch, 0)

			if date.Before(time.Now()) {
				continue
			}

			message := fmt.Sprintf("%s - %.0fC, Chance of rain: %.0f%%, %s\n", date.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

			if hour.ChanceOfRain > 40 {
				color.Cyan(message)
			} else if hour.TempC > 35 {
				color.Red(message)
			} else {
				fmt.Print(message)
			}
		}

	},
}

func init() {
	// https://www.youtube.com/watch?v=zPYjfgxYO7k

}
