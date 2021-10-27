// Package weather is doing Forecast.
package weather

// CurrentCondition is city's condition.
var CurrentCondition string

// CurrentLocation is the city name.
var CurrentLocation string

// Forecast is calculate the weather condition of city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
