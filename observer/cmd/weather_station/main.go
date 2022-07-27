package main

import (
	"github.com/rstliz/weather_observer"
)

func main() {
	weatherData :=
		weather_observer.NewWeatherData()

	currentDisplay := weather_observer.NewCurrentConditionsDisplay(weatherData)
	statisticsDisplay := weather_observer.NewStatisticsDisplay(weatherData)
	forecastDisplay := weather_observer.NewForecastDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)

	weatherData.RemoveObserver(currentDisplay)
	weatherData.RemoveObserver(statisticsDisplay)
	weatherData.RemoveObserver(forecastDisplay)
}
