package weather_observer

import "fmt"

type CurrentConditionsDisplay struct {
	temparature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	d := &CurrentConditionsDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(d)
	return d
}

func (d CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	d.temparature = temp
	d.humidity = humidity
	d.Display()
}

func (d CurrentConditionsDisplay) Display() {
	fmt.Printf("現在の気象状況: 温度%.2fF 湿度%.2f%%\n", d.temparature, d.humidity)
}
