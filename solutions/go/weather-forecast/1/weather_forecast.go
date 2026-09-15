// Package weather provides tools to track and broadcast the weather conditions
// of various cities across the globe. It allows Goblinocus meteorologists to 
// maintain a single, central record of the latest observed weather state.
package weather

var (
	// CurrentCondition stores the most recently reported weather state (such as "Sunny", 
	// "Rainy", or "Stormy"). Package users can read this variable to see the last 
	// known weather type, or manually update it when a new climate status is observed.
	CurrentCondition string

	// CurrentLocation stores the name of the city that was last checked for a weather report. 
	// Package users can read this variable to verify which region the current weather data 
	// belongs to, ensuring Goblinocus always tracks the correct territory.
	CurrentLocation  string
)

// Forecast creates and returns a formatted weather report string for a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
