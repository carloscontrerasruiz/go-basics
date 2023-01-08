package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
	fmt.Println("FIN DEL PROGRAMA")
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Inicia llamado %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Termina llamado %d\n", i)
}
