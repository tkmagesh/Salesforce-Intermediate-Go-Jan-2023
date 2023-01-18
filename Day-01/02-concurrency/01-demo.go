package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the execution using the go scheduler
	f2()

	// DO NOT DO THIS

	// time.Sleep(1 * time.Second) // blocking, and there by giving the opportunity for go scheduler to execute the scheduled goroutines

	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
