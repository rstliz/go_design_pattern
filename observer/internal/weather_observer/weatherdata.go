package weather_observer

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	d := &WeatherData{
		observers: make([]Observer, 0),
	}
	return d
}

func (d *WeatherData) RegisterObserver(o Observer) {
	d.observers = append(d.observers, o)
}

func (d *WeatherData) RemoveObserver(o Observer) {

	for i, observer := range d.observers {
		if observer == o {
			//d.observers = append(d.observers[:i], d.observers[i+1:]...)
			d.observers = d.observers[:i+copy(d.observers[i:], d.observers[i+1:])]
			break
		}
	}
}

func (d *WeatherData) NotifyObservers() {
	for _, observer := range d.observers {
		observer.Update(d.temperature, d.humidity, d.pressure)
	}
}

func (d *WeatherData) MeasurementsChanged() {
	d.NotifyObservers()
}

func (d *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	d.temperature = temperature
	d.humidity = humidity
	d.pressure = pressure
	d.MeasurementsChanged()
}

func (d WeatherData) GetTemperature(o Observer) {
}

func (d WeatherData) GetHumidity(o Observer) {
}

func (d WeatherData) GetPressure(o Observer) {
}

func (d WeatherData) GetObservers() []Observer {
	return d.observers
}
