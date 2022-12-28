package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	v := 1
	x = 8
	fmt.Println(v + x)

	m := make(map[string]int)
	m["Key"] = 6

	s := []int{1, 2, 3}

	s = append(s, 9)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// c := make(chan int)
	// go doSomething(c)
	// number := <-c
	// fmt.Println(number)

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("DONE")
	c <- 1
}
