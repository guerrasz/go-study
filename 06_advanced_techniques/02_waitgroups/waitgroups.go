package main

import (
	"fmt"
	"sync"
	"time"
)

// x = numWorkers
// y = numJobs
// creates x workers for y jobs all the workers loop over the tasks channel waiting for tasks
// those tasks are dinamically distributed among the x workers
func main() {
	// create wg
	var wg sync.WaitGroup

	// define num of jobs and workers
	numWorkers := 3
	numJobs := 10

	// set wg counter to numWorkers as it will wait for all workers to finish
	// workers only finish when tasks channel is closed and all tasks are done
	wg.Add(numWorkers)

	// create results and tasks channels with buffer size numJobs to avoid blocking
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	// start numWorkers workers as goroutines
	for i := range numWorkers {
		go worker(i+1, tasks, results, &wg)
	}

	// send numJobs tasks to tasks channel (here they are int from 1 to numJobs but could be any type)
	for i := range numJobs {
		tasks <- i + 1
	}
	// close tasks channel as no more tasks will be sent
	close(tasks)

	// goroutine to close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// collect results in realtime using range over results channel as it will be closed when done so no errors occur
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// worker that simulates doing some work and sends result to results channel and signals when done with wg.Done()
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	// defer wg.Done to signal when worker is done
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// simulate work time
	time.Sleep(time.Second)
	// loop over tasks channel to get tasks until channel is closed
	// this behavior happens in all workers dynamically so tasks are distributed among them
	for task := range tasks {
		results <- task * 2
	}
	fmt.Printf("Worker %d done\n", id)
}
