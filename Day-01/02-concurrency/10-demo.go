/*
Channel Behaviors:
	- A RECEIVE operation is ALWAYS a blocking operation
	- A SEND operation is blocked UNTIL a RECEIVE operation is initiated (conditions apply)

*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	//Approach-1
	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	//Approach-2
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		data := <-ch
		fmt.Println(data)
		wg.Done()
	}()
	ch <- 100
	wg.Wait()
}
