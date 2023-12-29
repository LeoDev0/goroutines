package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	concurrencyWorkers := 5
	in := make(chan int)
	done := make(chan struct{}) // Use an empty struct instead of []byte for signaling completion
	taskSize := 200

	var wg sync.WaitGroup
	wg.Add(concurrencyWorkers) // Set the WaitGroup counter to the number of workers

	startTime := time.Now()

	go func() {
		defer close(in)
		for i := 0; i <= taskSize; i++ {
			in <- i
		}
	}()

	for x := 0; x < concurrencyWorkers; x++ {
		go ProcessWorker(&wg, in, x)
	}

	wg.Wait() // Wait for all workers to finish

	// Signal the main function that all tasks are done
	close(done)

	processingTime := time.Since(startTime)

	// Keep the main function alive for a short time to ensure the done channel is processed
	time.Sleep(1 * time.Second)

	fmt.Println("Total time to finish", taskSize, "processes:", processingTime.Truncate(time.Millisecond))
}

func ProcessWorker(wg *sync.WaitGroup, in chan int, worker int) {
	defer wg.Done() // Signal the WaitGroup that this worker has finished
	for x := range in {
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		fmt.Println("Worker", worker, ":", int(x))
	}
}
