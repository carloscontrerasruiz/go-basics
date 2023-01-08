package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 5) //buffered channels
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}

	wg.Wait()
	fmt.Println("FIN DEL PROGRAMA")
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Inicia llamado %d\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Termina llamado %d\n", i)
	<-c
}
