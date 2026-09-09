package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go showNumbers()
	fmt.Println("Main function running")

	time.Sleep(1 * time.Second)

	go worker("Worker A")
	go worker("Worker B")
	go worker("Worker C")

	time.Sleep(1 * time.Second)

	ch := make(chan int)
	go sendNumbers(ch)
	for num := range ch {
		fmt.Println(num)
	}

	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	numbers := []int{1, 2, 3, 4, 5}
	chn := make(chan int)
	for _, num := range numbers {
		go square(num, chn)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-chn)
	}

	jobs := make(chan int)
	for w := 1; w <= 3; w++ {
		go work(w, jobs)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	ch = make(chan int)
	go generator(1, ch)
	go generator(10, ch)
	go generator(100, ch)

	for i := 0; i < 9; i++ {
		fmt.Println(<-ch)
	}

	jobs = make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 10; r++ {
		fmt.Println("Result:", <-results)
	}

	channelA := make(chan string)
	channelB := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		channelA <- "Message from A"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channelB <- "Message from B"
	}()

	select {
	case msg := <-channelA:
		fmt.Println(msg)
	case msg := <-channelB:
		fmt.Println(msg)
	}

	//Exercise 10
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {

			mu.Lock()
			counter++
			mu.Unlock()

			wg.Done()

		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
