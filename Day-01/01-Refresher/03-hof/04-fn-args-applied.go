package main

import (
	"fmt"
	"log"
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
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)

}

func logOperation(operation func(int, int), x, y int) {
	log.Println("operation started")
	operation(x, y)
	log.Println("operation completed")
}

func logAdd(x, y int) {
	log.Println("operation started")
	add(x, y)
	log.Println("operation completed")
}

func logSubtract(x, y int) {
	log.Println("operation started")
	subtract(x, y)
	log.Println("operation completed")
}

func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
