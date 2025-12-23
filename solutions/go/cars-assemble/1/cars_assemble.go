package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var successfulCars float64 = (float64(productionRate) * successRate)/100
    return successfulCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    var successfulCars int = int(perHour) / 60
    return successfulCars
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var rest uint = uint(carsCount) % 10             // 27%10 == 7   30/10 == 0 
	var division uint = uint(carsCount) / 10        // 30/10 == 3  27/10 == 2 
    var totalCost uint = division * 95000 + rest * 10000
    return totalCost
}
