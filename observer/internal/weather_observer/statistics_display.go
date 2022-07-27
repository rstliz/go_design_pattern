package weather_observer

import "fmt"

type StatisticsDisplay struct {
	maxTemp     float64
	minTemp     float64
	tempSum     float64
	numReadings int
	weatherData Subject
}

func NewStatisticsDisplay(weatherData Subject) *StatisticsDisplay {
	d := &StatisticsDisplay{
		maxTemp:     0.0,
		minTemp:     200.0,
		tempSum:     0.0,
		numReadings: 0,
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(d)
	return d
}

func (d StatisticsDisplay) Update(temp float64, humidity float64, pressure float64) {

	d.numReadings++
	d.tempSum += temp
	if d.maxTemp < temp {
		d.maxTemp = temp
	}
	if d.minTemp > temp {
		d.minTemp = temp
	}
	d.Display()
}

func (d StatisticsDisplay) Display() {
	fmt.Printf("平均/最高/最低 気温 = %.2f/%.2f/%.2f\n", d.tempSum/float64(d.numReadings), d.maxTemp, d.minTemp)
}
