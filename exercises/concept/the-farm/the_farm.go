package thefarm

import "errors"
import "strconv"

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, c int) (float64, error) {

	totalAmountFodder, err := f.FodderAmount(c)

	if err != nil {
		return 0.0, err
	}

	factor, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return totalAmountFodder * factor / float64(c), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, c int) (float64, error) {
	if c <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(f, c)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(n int) error {
	cowStr := strconv.Itoa(n)

	if n < 0 {
		return errors.New(cowStr + " cows are invalid: there are no negative cows")
	} else if n == 0 {
		return errors.New(cowStr + " cows are invalid: no cows don't need food")
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
