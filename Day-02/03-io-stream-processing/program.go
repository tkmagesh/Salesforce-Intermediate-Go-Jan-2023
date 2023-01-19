package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	dataCh := make(chan int)
	wg.Add(1)
	go Source("data1.dat", dataCh, wg)

	wg.Add(1)
	go Source("data2.dat", dataCh, wg)

	evenCh, oddCh := Splitter(dataCh)
	evenSumCh := Sum(evenCh)
	oddSumCh := Sum(oddCh)

	doneCh := Merger(evenSumCh, oddSumCh)

	wg.Wait()
	close(dataCh)

	<-doneCh
	fmt.Println("Done")
}

func Merger(evenSumCh, oddSumCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create("result.txt")
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		for i := 0; i < 2; i++ {
			select {
			case evenSum := <-evenSumCh:
				file.WriteString(fmt.Sprintf("Even Total : %d\n", evenSum))
			case oddSum := <-oddSumCh:
				file.WriteString(fmt.Sprintf("Odd Total : %d\n", oddSum))
			}
		}
		// close(doneCh)
		doneCh <- struct{}{}
	}()
	return doneCh
}
func Sum(ch <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		result := 0
		for no := range ch {
			result += no
		}
		sumCh <- result
	}()
	return sumCh
}

func Splitter(dataCh chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		defer close(evenCh)
		defer close(oddCh)
		for no := range dataCh {
			if no%2 == 0 {
				evenCh <- no
			} else {
				oddCh <- no
			}
		}
	}()
	return evenCh, oddCh
}

func Source(fileName string, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if no, err := strconv.Atoi(txt); err == nil {
			dataCh <- no
		}
	}
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
