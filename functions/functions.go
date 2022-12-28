package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5

	y := func() int {
		return x * 2
	}

	z := func(number int) int {
		return number * 2
	}

	fmt.Println(y())
	fmt.Println(z(5))

	funcAsParams(z)

	c := make(chan int)
	go func() {
		fmt.Println("Starting this function")
		time.Sleep(5 * time.Second)
		c <- 1
	}()

	resilt := <-c
	fmt.Println(resilt)
}

func funcAsParams(f func(number int) int) {
	number := f(1090)
	fmt.Println("Desde otra funcion")
	fmt.Println(number)
}
