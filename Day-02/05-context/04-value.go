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
	valCtx := context.WithValue(rootCtx, "root-key", "root-val")
	wg.Add(1)
	// cancelCtx, cancel := context.WithCancel(rootCtx)
	cancelCtx, cancel := context.WithCancel(valCtx)
	defer cancel()

	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()

	go fn(wg, cancelCtx)
	wg.Wait()
	fmt.Println("Thank you!")
}

func fn(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	fmt.Println("[fn] root-key =", ctx.Value("root-key"))
	i := 0

	wg.Add(1)
	go f1(wg, ctx)

	wg.Add(1)
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	go f2(wg, timeoutCtx)

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fn - Cancel signal received.. shutting down")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			i++
			fmt.Printf("fn invoked - %d\n", i)
		}
	}
	fmt.Println("fn Done")
}

func f1(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	i := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f1 - Cancel signal received.. shutting down")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			i++
			fmt.Printf("f1 invoked - %d\n", i*10)
		}
	}
	fmt.Println("f1 Done")
}

func f2(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	fmt.Println("[f2] root-key =", ctx.Value("root-key"))
	i := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("f2 - Cancel signal received.. shutting down")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			i++
			fmt.Printf("f2 invoked - %d\n", i*3)
		}
	}
	fmt.Println("f2 Done")
}
