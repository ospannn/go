package main

import (
	"fmt"

	"example.com/mathlib"
)

func main() {
	fmt.Println("Қосынды:", mathlib.Add(5, 3))
	fmt.Println("Көбейтінді:", mathlib.Multiply(5, 3))
}
