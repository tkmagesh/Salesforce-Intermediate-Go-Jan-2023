package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	wg.Add(1)
	cancelCtx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel()

	go fn(wg, cancelCtx)
	wg.Wait()
	fmt.Println("Thank you!")
}

func fn(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	i := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel signal received.. shutting down")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			i++
			fmt.Printf("fn invoked - %d\n", i)
		}
	}
	fmt.Println("Done")
}
