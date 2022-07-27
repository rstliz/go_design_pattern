```mermaid
classDiagram

class Subject {
  <<interface>>
  +RegisterObserver()
  +RemoveObserver()
  +NotifyObserver()
}

class Observer {
  <<interface>>
  +Update()
}

class DisplayElement {
  <<interface>>
  +Display()
}

class WeatherData {
  +RegisterObserver()
  +RemoveObserver()
  +NotifyObserver()

  +GetTemperature()
  +GetHumidity()
  +GetPressure()
  +MeasurementsChanged()
}

class CurrentConditionsDisplay {
  +Update()
  +Display()
}

class StatisticsDisplay {
  +Update()
  +Display()
}

class ThirdPartyDisplay {
  +Update()
  +Display()
}

class ForecastDisplay {
  +Update()
  +Display()
}


Subject --> Observer : observers
Subject <|.. WeatherData
WeatherData <-- CurrentConditionsDisplay : subject

CurrentConditionsDisplay ..|> DisplayElement
CurrentConditionsDisplay ..|> Observer

StatisticsDisplay ..|> DisplayElement
StatisticsDisplay ..|> Observer

ThirdPartyDisplay ..|> DisplayElement
ThirdPartyDisplay ..|> Observer

ForecastDisplay ..|> DisplayElement
ForecastDisplay ..|> Observer

```
