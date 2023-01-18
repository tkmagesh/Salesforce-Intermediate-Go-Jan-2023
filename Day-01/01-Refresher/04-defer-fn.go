package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	deferred - main")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	deferred - f1[1]")
		}()
		defer func() {
			fmt.Println("	deferred - f1[2]")
		}()
		defer func() {
			fmt.Println("	deferred - f1[3]")
		}()
	*/

	defer fmt.Println("	deferred - f1[1]")
	defer fmt.Println("	deferred - f1[2]")
	defer fmt.Println("	deferred - f1[3]")

	fmt.Println("f1 started")
	f2Result := f2()
	fmt.Println("f2Result :", f2Result)
	fmt.Println("f1 completed")
}

func f2() (x int) {
	defer func() {
		fmt.Println("	deferred - f2")
		x = 300
	}()
	fmt.Println("f2 started")
	// panic("dummy panic")
	fmt.Println("f2 completed")
	x = 100
	return
}
