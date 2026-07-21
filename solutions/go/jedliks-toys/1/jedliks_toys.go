package jedlik

import (
	"strconv"
)

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery = c.battery - c.batteryDrain
		c.distance = c.distance + c.speed
	} 
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	res := strconv.Itoa(c.distance)
	return "Driven " + res + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	res := strconv.Itoa(c.battery)
	return "Battery at " + res + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	if c.speed*c.battery >= c.batteryDrain*trackDistance {
		return true
	}
	return false
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
