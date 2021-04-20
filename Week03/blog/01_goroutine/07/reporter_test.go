package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestReporter_shutdown_datarace(t *testing.T) {
	r := NewReporter(1, 10)
	var wg sync.WaitGroup

	stop := make(chan struct{})
	go func() {
		r.Run(stop)
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				time.Sleep(time.Microsecond)
				r.Report(fmt.Sprintf("test: %d", i))
			}
			fmt.Println("done")
		}()
	}

	time.AfterFunc(time.Microsecond, func() {
		fmt.Println("will stop...")
		close(stop)
	})
	wg.Wait()
}
