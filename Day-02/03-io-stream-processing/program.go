package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sum1 := test("data1.dat")
	sum2 := test("data2.dat")
	fmt.Println(sum1 + sum2)
}

/* How to read from file and convert text to int */
func test(fileName string) int {
	total := 0
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if no, err := strconv.Atoi(txt); err == nil {
			total += no
		}
	}
	return total
}
