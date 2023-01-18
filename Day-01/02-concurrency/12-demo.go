/* sending and receiving multiple values using a channel */

package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	count := 10
	go genNos(ch, count)
	for i := 1; i <= count; i++ {
		fmt.Println(<-ch)
	}
}

//producer
func genNos(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
}
