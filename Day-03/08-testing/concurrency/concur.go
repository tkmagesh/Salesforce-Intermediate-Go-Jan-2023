package concurrency

import "sync"

type Counter struct {
	sync.Mutex
	No int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.No++
	}
	c.Unlock()
}
