// Package weather provides weather forecast.
package weather

var (
	//CurrentCondition represents the actual weather condition.
	CurrentCondition string
	//CurrentLocation is the name of the town for which forecast are made.
	CurrentLocation  string
)

// Forecast takes in input the city and the actual weather condition and return a string the weather condition for that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
