package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	slots chan struct{}
}

func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		slots: make(chan struct{}, max),
	}
}

func (s *Semaphore) Acquire() {
	s.slots <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.slots
}

func (s *Semaphore) TryAcquire(id int) bool {
	select {
	case s.slots <- struct{}{}:
		fmt.Printf("Горутина %d начала работу\n", id)
		return true
	default:
		fmt.Println(id, false)
		return false
	}
}

func main() {
	sem := NewSemaphore(3)

	for i := 0; i < 10; i++ {
		go func(id int) {
			sem.Acquire()
			defer sem.Release()
			fmt.Printf("Горутина %d начала работу\n", id)
			time.Sleep(3 * time.Second)
		}(i)
	}
	time.Sleep(10 * time.Second)
}
