package main

import (
	"fmt"
	"math/rand"
	"time"
	"worker-demo/worker"
)

type MyWork struct {
	Id int
}

func (myWork MyWork) Task() {
	fmt.Printf("task [%d] started...\n", myWork.Id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Printf("task [%d] completed...\n", myWork.Id)
}

func main() {
	w := worker.New(5)
	for i := 1; i <= 20; i++ {
		w.Add(MyWork{Id: i})
	}
	fmt.Println("All tasks are assigned....")
	fmt.Scanln()
	w.Shutdown() //wait for all the tasks to complete
}
