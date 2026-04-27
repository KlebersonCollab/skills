package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID int
}

// WorkerPool processes tasks concurrently using a fixed number of workers.
func WorkerPool(ctx context.Context, tasks []Task, numWorkers int) {
	jobs := make(chan Task, len(tasks))
	results := make(chan int, len(tasks))
	var wg sync.WaitGroup

	// Start workers.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						return
					}
					fmt.Printf("Worker %d processing task %d\n", workerID, job.ID)
					time.Sleep(100 * time.Millisecond) // Simulate work.
					results <- job.ID
				case <-ctx.Done():
					fmt.Printf("Worker %d shutting down: %v\n", workerID, ctx.Err())
					return
				}
			}
		}(w)
	}

	// Send jobs.
	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	// Wait for workers to finish.
	wg.Wait()
	close(results)

	fmt.Println("All tasks processed.")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tasks := []Task{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}}
	WorkerPool(ctx, tasks, 3)
}
