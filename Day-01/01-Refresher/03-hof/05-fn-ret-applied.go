package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		log.Println("operation started")
		add(100, 200)
		log.Println("operation completed")

		log.Println("operation started")
		subtract(100, 200)
		log.Println("operation completed")
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	//function composition for "logging"
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	/* function composition for "profiling" */
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	/* combining logging and profiling */
	/*
		logAdd := getLogOperation(add)
		profileLogAdd := getProfileOperation(logAdd)
	*/
	profileLogAdd := getProfileOperation(getLogOperation(add))
	profileLogAdd(100, 200)

	profileLogSubtract := getProfileOperation(getLogOperation(subtract))
	profileLogSubtract(100, 200)

}

func getProfileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

/*
func logOperation(operation func(int, int), x, y int) {
	log.Println("operation started")
	operation(x, y)
	log.Println("operation completed")
}
*/

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("operation started")
		operation(x, y)
		log.Println("operation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
