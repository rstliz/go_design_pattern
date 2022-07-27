package weather_observer

import "fmt"

type ForecastDisplay struct {
	currentPressure float64
	lastPressure    float64
	weatherData     Subject
}

func NewForecastDisplay(weatherData Subject) *ForecastDisplay {
	d := &ForecastDisplay{
		currentPressure: 29.92,
		weatherData:     weatherData,
	}
	weatherData.RegisterObserver(d)
	return d
}

func (d ForecastDisplay) Update(temp float64, humidity float64, pressure float64) {
	d.lastPressure = d.currentPressure
	d.currentPressure = pressure
	d.Display()
}

func (d ForecastDisplay) Display() {
	fmt.Print("予報: ")
	if d.currentPressure > d.lastPressure {
		fmt.Println("天候は良くなります！")
	} else if d.currentPressure == d.lastPressure {
		fmt.Println("ほとんど同じです。")
	} else if d.currentPressure < d.lastPressure {
		fmt.Println("より寒く雨模様の天候に注意してください。")
	}
}
