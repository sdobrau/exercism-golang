// Package weather provides forecast information.
package weather

var (
    // CurrentCondition is the current condition of weather forecast.
	CurrentCondition string
    // CurrentLocation is the current location of the weather forecast.
	CurrentLocation  string 
)

// Forecast takes a city and condition as input and returns a weather report for input city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
