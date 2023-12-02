package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var sum int64 = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64 = 0
	for _,v := range m {
		sum += v
	}
	return sum
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _,v := range m {
		sum += v
	}
	return sum
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _,v := range m {
		sum += v
	}
	return sum
}

func main() {
	ints := map[string]int64 {
		"first" : 35,
		"second" : 12,
	}

	floats := map[string]float64 {
		"first": 12.123,
		"second": 44.12331,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",SumInts(ints),SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n",SumIntsOrFloats[string, int64](ints),SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums type inferred: %v and %v\n",SumIntsOrFloats(ints),SumIntsOrFloats(floats))
	fmt.Printf("Constraint Sums: %v and %v\n",SumNumbers(ints),SumNumbers(floats))
}