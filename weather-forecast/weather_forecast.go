// Package weather provides a weather forecasting tools.
package weather

// CurrentCondition tells us current weather condition.
var CurrentCondition string
// CurrentLocation is the worl location for forecasting.
var CurrentLocation string

// Forecast returns a string that tells us weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
