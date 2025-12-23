package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64((float64(productionRate) * successRate)/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    rest := uint(carsCount) % 10              
	division := uint(carsCount) / 10        
    return uint(division * 95000 + rest * 10000)
}
