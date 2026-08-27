package thefarm

import (
    "fmt"
    "errors"
)


// TODO: define the 'DivideFood' function
func DivideFood(fodder FodderCalculator, cows int) (float64, error) {

	if cows <= 0 {
		return 0, errors.New("something went wrong")
	}

	countAll, err := fodder.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	koeff, err := fodder.FatteningFactor()
	if err != nil {
		return 0, err
	}

	countOne := countAll / float64(cows)
	return countOne * koeff, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodder FodderCalculator, cows int) (float64, error) {
if cows > 0 {
	res, err := DivideFood(fodder, cows)
	if err != nil {
		return 0, err
	}
	return res, nil
}
return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	CountCow int
	message  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", e.CountCow, e.message)
}

func ValidateNumberOfCows(count int) error {
	if count < 0 {
		return &InvalidCowsError{CountCow: count, message: "there are no negative cows"}
	}

	if count == 0 {
		return &InvalidCowsError{CountCow: count, message: "no cows don't need food"}
	}

	return nil

}


