package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	var d int
	for d == 0 {
		d = rand.Intn(20)
		if d != 0 {
			return d
		}
	}
	return 0
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	var d float64
	for 0.0 <= d && d < 12.0 {
		d = rand.Float64() * 12
		return d
	}
	return 0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animalsFirstSlice := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animalsFirstSlice), func(i, j int) {
		animalsFirstSlice[i], animalsFirstSlice[j] = animalsFirstSlice[j], animalsFirstSlice[i]
	})
	return animalsFirstSlice
}
