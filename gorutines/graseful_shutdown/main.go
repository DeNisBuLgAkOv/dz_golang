package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var numGoroutines int
	flag.IntVar(&numGoroutines, "n", 1, "")
	flag.Parse()

	stopCh := make(chan os.Signal, 1)
	resultCh := make(chan float64, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go calculatePi(i, numGoroutines, resultCh, stopCh)
	}

	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	close(stopCh)
	fmt.Println("Безопасное завершение")

	totalSum := 0.0
	close(resultCh)

	for partialSum := range resultCh {
		totalSum += partialSum
	}

	fmt.Println(totalSum * 4)
}

func calculatePi(goroutineID int, numGoroutines int, resultCh chan float64, stopCh chan os.Signal) {

	localSum := 0.0
	for i := goroutineID; ; i += numGoroutines {

		select {
		case <-stopCh:
			resultCh <- localSum
			return
		default:
			term := 1.0 / float64(2*i+1)
			if i%2 == 1 {
				term = -term
			}
			localSum += term
		}

	}
}
