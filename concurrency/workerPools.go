package main

import "fmt"

func Fibonnacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnacci(n-1) + Fibonnacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib for %d\n", id, job)
		fib := Fibonnacci(job)
		fmt.Printf("Worker with id %d started fib for %d and result %d\n", id, job, fib)
		results <- fib
	}
}

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, v := range tasks {
		jobs <- v
	}
	close(jobs)

	for i := 0; 1 < len(tasks); i++ {
		<-results
	}
}
