package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 10
	var demonimator int = 2
	var result, remainder, err = intDivision(numerator, demonimator)

	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	// USING SWITCH STATEMENT

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

}

func intDivision(numerator int, demonimator int) (int, int, error) {
	var err error
	if demonimator == 0 {
		err = errors.New("Cannot Divide By Zero")
		return 0, 0, err
	}
	var result int = numerator / demonimator
	var remainder int = numerator % demonimator

	return result, remainder, err
}
