package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Task struct {
	value int
	err   error
}

func HandleTasks(value int) Task {
	if value == 3 || value == 8 || value == 12 {
		return Task{value: value, err: errors.New("ошибка")}
	}
	return Task{value: value, err: nil}
}

func Worker(id int, tasksCh <-chan Task, resultCh chan<- int, counterErr *int64, maxErrors int64) {

	for task := range tasksCh {
		if atomic.LoadInt64(counterErr) == maxErrors {
			return
		}

		if task.err != nil {
			atomic.AddInt64(counterErr, 1)
			continue
		}
		resultCh <- task.value
	}
}

func main() {
	var (
		wg         sync.WaitGroup
		numWorkers       = 3
		numTasks         = 20
		maxErrors  int64 = 1
		counterErr int64 = 0
	)

	tasksCh := make(chan Task, numTasks)
	resultsCh := make(chan int, numTasks)

	go func() {
		for i := 1; i <= numTasks; i++ {
			tasksCh <- HandleTasks(i)
		}
		close(tasksCh)
	}()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			Worker(workerID, tasksCh, resultsCh, &counterErr, maxErrors)
		}(i)
	}
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	for res := range resultsCh {
		fmt.Println("Обработано:", res)
	}
}
