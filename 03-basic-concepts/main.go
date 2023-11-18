package main

import (
	"fmt"
	"03-basic-concepts/step1"
)

func main() {
	fmt.Printf("Max int32: %v Overflow: %v\n", step1.Sum(1), step1.Sum(2))
}
