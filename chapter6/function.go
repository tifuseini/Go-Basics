package main

import "fmt"

func main() {

	xs := []float64{90, 89, 34, 34, 45}
	var averaged = average(xs)
	fmt.Println(averaged)
	fmt.Println("89")
	fmt.Println("Happy New Year")

}

func average(xs []float64) float64 {

	total := 0.0

	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))

}
