package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(sloi []string, timeDoneMin int) int {
	if timeDoneMin == 0 {
		timeDoneMin = 2
	}
	return len(sloi) * timeDoneMin
}

// TODO: define the 'Quantities()' function
func Quantities(sloi []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, v := range sloi {
		if v == "noodles" {
			noodles++
		}
		if v == "sauce" {
			sauce++
		}
	}

	return noodles * 50, sauce * 0.2

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend, my []string) []string {
	lenMy := len(my)
	lenFri := len(friend)

	my[lenMy-1] = friend[lenFri-1]

	return my

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(needTwo []float64, value int) []float64 {
	var slice []float64
	slice = make([]float64, len(needTwo))
	copy(slice, needTwo)

	for i, v := range slice {
		slice[i] = v * (float64(value) / 2)
	}
	return slice

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
