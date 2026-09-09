package main

import (
	"fmt"
	"time"
)

func showNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func worker(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, "running")
		time.Sleep(200 * time.Millisecond)
	}
}

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func sendNumber(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
}

func square(n int, chn chan int) {
	chn <- n * n
}

func work(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d processed job %d/n", id, job)
	}
}

func generator(start int, ch chan int) {
	for i := start; i < start+3; i++ {
		ch <- 1

	}
}

func workerPool(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d/n", id, job)
		results <- job * 2
	}
}
