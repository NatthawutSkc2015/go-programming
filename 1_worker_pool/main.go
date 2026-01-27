package main

import (
	"fmt"
	"sync"
	"time"
)

// RunWorkers creates a worker pool to process jobs concurrently
func RunWorkers(numWorkers, numJobs int) {
	// Create channels for jobs and done signal
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Worker processes jobs from the channel
			for jobID := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, jobID)
				time.Sleep(1 * time.Second) // Simulate work
			}
		}(w)
	}

	// Send jobs to channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close channel when all jobs are sent

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs completed!")
}

func main() {
	fmt.Println("=== Worker Pool Example ===")
	RunWorkers(3, 10) // 3 workers, 10 jobs
}
