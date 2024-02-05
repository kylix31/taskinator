package taskinator

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Task represents a scheduled task
type Task struct {
	Name       string
	Schedule   time.Duration
	MaxRetries int
	RetryDelay time.Duration
	Action     func() error
}

// Scheduler represents a task scheduler
type Scheduler struct {
	tasks []Task
}

// NewScheduler creates a new task scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{tasks: make([]Task, 0)}
}

// AddTask adds a task to the scheduler
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Start starts the task scheduler for the specified duration
func (s *Scheduler) Start(duration ...time.Duration) {
	var wg sync.WaitGroup

	var endTime time.Time
	var hasDuration bool

	// Check if a duration was provided
	if len(duration) > 0 && duration[0] > 0 {
		hasDuration = true
		endTime = time.Now().Add(duration[0])
	}

	for _, task := range s.tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			ticker := time.NewTicker(t.Schedule)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					if err := executeWithRetry(t); err != nil {
						log.Printf("Task %s failed after retrying: %v\n", t.Name, err)
					}
				}

				// Check if the scheduler has run for the specified duration
				if hasDuration && time.Now().After(endTime) {
					return
				}
			}
		}(task)
	}

	fmt.Println("Task scheduler started...")

	wg.Wait()
}

// executeWithRetry executes a task with retry logic
func executeWithRetry(task Task) error {
	for attempt := 0; attempt <= task.MaxRetries; attempt++ {
		if err := task.Action(); err == nil {
			return nil // Task executed successfully, no need for retry
		} else {
			fmt.Printf("Task %s failed (Attempt %d), retrying...\n", task.Name, attempt+1)
			time.Sleep(task.RetryDelay) // Delay before retrying
		}
	}
	return fmt.Errorf("Task %s exceeded maximum retries", task.Name)
}
