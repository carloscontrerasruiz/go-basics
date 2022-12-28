package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sumArgsVar(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func values(x int) (int, int, int) {
	return x * 2, x * 3, x * 4
}

func namedValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(10, 45))

	fmt.Println(sumArgsVar(10, 45, 7, 8, 4))

	fmt.Println(values(10))

	fmt.Println(namedValues(1000))
}
