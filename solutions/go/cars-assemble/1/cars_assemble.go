package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	res := float64(productionRate) * (successRate / 100)
	return res
}


// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	res := (float64(productionRate) * (successRate) / 100) / 60
	return int(res)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	resSale := (carsCount / 10) * 95_000
	resNotSale := carsCount % 10 * 10_000
	return uint(resSale + resNotSale)
}
