package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	or := func(channels ...<-chan interface{}) <-chan struct{} {
		ctx, cansel := context.WithCancel(context.Background())
		fmt.Println(ctx)

		for _, ch := range channels {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					cansel()
				case <-ctx.Done():
					return
				}
			}(ch)
		}
		fmt.Println(ctx.Done())

		return ctx.Done()
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
