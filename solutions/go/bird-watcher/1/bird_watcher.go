package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var res int
	for _, v := range birdsPerDay {
		res += v
	}
	return res 
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var res int
	switch week {
	case 1:
		for i := 0; i <= 6; i++ {
			res += birdsPerDay[i]
		}
	case 2:
		for i := 7; i <= 13; i++ {
			res += birdsPerDay[i]
		}
	case 3:
		for i := 14; i <= 20; i++ {
			res += birdsPerDay[i]
		}
	case 4:
		for i := 21; i <= 28; i++ {
			res += birdsPerDay[i]
		}
	}
	return res
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay

}
