package main

import "fmt"

type NumberTypeIntOrFloat interface {
	int | float64
}

func sumGeneric[V int | float64](nums []V) V {
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumGenericWithInterface[V NumberTypeIntOrFloat](nums []V) V {
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumInt(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumFloat(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	sliceInt := []int{1, 2, 3, 4}
	sliceFloat := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(sumInt(sliceInt))
	fmt.Println(sumFloat(sliceFloat))
	fmt.Println(sumGeneric(sliceInt))
	fmt.Println(sumGeneric(sliceFloat))
	fmt.Println(sumGenericWithInterface(sliceInt))
	fmt.Println(sumGenericWithInterface(sliceFloat))
}
