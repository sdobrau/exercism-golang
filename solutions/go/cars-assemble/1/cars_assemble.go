package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := (CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
    return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	numberOfTens := carsCount / 10
    remainderOfTens := carsCount % 10
    costOfTens := 95000 * numberOfTens
    costOfRemainders := 10000 * remainderOfTens
    totalCost := costOfTens + costOfRemainders
    return uint(totalCost)
}
