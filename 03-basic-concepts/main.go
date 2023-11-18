package main

import (
	"fmt"
	"03-basic-concepts/step1"
)

func main() {
	fmt.Printf("Max int32: %v Overflow: %v\n", step1.Sum(1), step1.Sum(2))

	fmt.Printf("Distance between Yerevan and Tolyatti: %v km\n", step1.GetDistanceFromLatLonInKm(40.1534924, 44.4061668, 53.5219693,49.3215277))
}
