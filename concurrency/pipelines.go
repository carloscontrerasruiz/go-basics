package main

import "fmt"

func Generator(c chan<- int) { //operador flecha a la derecha es de escritura
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

func Double(in <-chan int, out chan<- int) { //operador flecha a la izquierda es de lectura
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
