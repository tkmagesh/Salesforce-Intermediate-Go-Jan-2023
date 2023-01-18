package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rand.Seed(7)
	count := flag.Int("count", 0, "Number of goroutines to spin up")
	flag.Parse()
	fmt.Printf("Starting %d goroutines. Hit ENTER to start....\n", *count)
	fmt.Scanln()
	for i := 1; i <= *count; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Printf("All goroutines completed. Hit ENTER to shutdown....")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}

func f2() {
	fmt.Println("f2 invoked")
}
