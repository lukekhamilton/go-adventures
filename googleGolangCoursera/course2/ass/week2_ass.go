package main

import (
	"fmt"
	"math"
	"strconv"
)

// s =½ a t2 + vt + s
var acceleration float64        // a
var initialVelocity float64     // v
var initialDisplacement float64 // s
var time float64                // t

func inputData(msg string) float64 {
	validInput := false
	var input float64
	for validInput != true {
		fmt.Print(msg)
		var a string
		fmt.Scan(&a)

		input, err := strconv.ParseFloat(a, 64)

		if err != nil {
			fmt.Println("Sorry invalid input. Error message:")
			fmt.Println(err)
			fmt.Println("Please try again")
		} else {
			validInput = true
			return input
		}
	}
	return input
}

func captureInput(acceleration *float64, initialVelocity *float64, initialDisplacement *float64, time *float64) {
	*acceleration = inputData("Enter acceleration: ")
	*initialVelocity = inputData("Enter initial velocity: ")
	*initialDisplacement = inputData("Enter initial displacement: ")
	*time = inputData("Enter travel time: ")
}

func main() {
	captureInput(&acceleration, &initialVelocity, &initialDisplacement, &time)
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(fn(time))
}

// GenDisplaceFn ...
// a acceleration
// v initialVelocity
// s initialDisplacement
func GenDisplaceFn(a float64, v float64, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		// s =½ a t2 + vt + s
		return .5*a*(math.Pow(t, 2)) + (v * t) + s
	}
}
