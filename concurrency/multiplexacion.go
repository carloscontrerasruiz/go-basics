package main

import (
	"fmt"
	"time"
)

func doSomething(i time.Duration, c chan<- int, param int) {
	fmt.Printf("Inicia llamado %d\n", param)
	time.Sleep(i)
	fmt.Printf("Termina llamado %d\n", param)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	//No imprta que el channel 2 termine antes el programa se bloquea
	// hasta que e channel 1 termina por lo cual el channel 2 esta en espera
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}
