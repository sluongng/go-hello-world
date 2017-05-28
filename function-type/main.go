package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	return !isOdd(integer)
}

func isZero(integer int) bool {
	if integer == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{0, 1, 2, 3, 14, 28, 99, 131}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)

	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)

	nonZero := filter(slice, func(x int) bool {
		return !isZero(x)
	})
	fmt.Println("Non zero elements are: ", nonZero)
}