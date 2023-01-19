package main

import "fmt"

func main() {
	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	data := <-ch
	fmt.Println(data)
	fmt.Println(len(ch))

	data = <-ch
	fmt.Println(data)
	fmt.Println(len(ch))

}
