package main

import (
	"fmt"
	"time"
)

func main() {
	// define number of workers and jobs
	numWorkers := 3
	numJobs := 10

	// create channels for tasks and results
	tasks := make(chan int, numJobs)

	results := make(chan int, numJobs)

	// start numWorkers worker goroutines
	for i := range numWorkers {
		go worker((i + 1), tasks, results)
	}

	// send tasks
	for i := range numJobs {
		tasks <- i
	}

	// close tasks channel, no more tasks will be sent but workers can finish processing existing ones
	close(tasks)

	// collect results, loop over results channel numJobs times to get all results
	for range numJobs {
		fmt.Printf("Result: %d\n", <-results)
	}
	close(results)

}

// simulates a worker that processes tasks from the tasks channel and sends results to the results channel
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second)
		results <- task * 2
	}
}
