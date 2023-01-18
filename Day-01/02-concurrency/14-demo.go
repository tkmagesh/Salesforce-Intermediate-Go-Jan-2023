/* sending and receiving multiple values using a channel */

package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	dataCh := genNos()
	for {
		if data, isOpen := <-dataCh; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
}

//producer
func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		count := 5
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
