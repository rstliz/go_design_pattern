module cmd_weather_station

go 1.18

require (
	github.com/rstliz/weather_observer v0.0.1
)

replace github.com/rstliz/weather_observer => ../../internal/weather_observer
