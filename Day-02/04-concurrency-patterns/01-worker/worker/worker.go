package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	workQueue chan Work
	wg        sync.WaitGroup
}

func (w *Worker) Add(wk Work) {
	w.workQueue <- wk
}

func (w *Worker) Shutdown() {
	close(w.workQueue)
	w.wg.Wait()
}

func New(count int) *Worker {
	w := &Worker{
		workQueue: make(chan Work),
	}
	for i := 1; i <= count; i++ {
		w.wg.Add(1)
		go func(id int) {
			fmt.Printf("Worker [%d] started...\n", id)
			for wk := range w.workQueue {
				wk.Task()
			}
			fmt.Printf("Worker [%d] shutdown...\n", id)
			w.wg.Done()
		}(i)
	}

	return w
}
