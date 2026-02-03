package main

import (
	"fmt"
	"sync"
	"time"
)

// RunWorkers creates a worker pool to process jobs concurrently
func RunWorkers(numWorkers, numJobs int) {
	// Create channels for jobs and results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Worker processes jobs from the channel
			for jobID := range jobs {
				time.Sleep(1 * time.Second) // Simulate work
				// Send result (jobID * 2) to results channel
				results <- jobID * 2
			}
		}(w)
	}

	// Start a goroutine to close results channel after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Send jobs to channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close channel when all jobs are sent

	// Wait for and collect all results
	fmt.Println("Collecting results:")
	for result := range results {
		fmt.Printf("Result received: %d\n", result)
	}

	fmt.Println("All jobs completed!")
}

func main() {
	fmt.Println("=== Worker Pool Example ===")
	RunWorkers(3, 10) // 3 workers, 10 jobs
}
