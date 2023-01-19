package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	closed    bool
	sync.Mutex
}

var ErrorPoolClosed = errors.New("pool closed")

func (p *Pool) Acquire() (io.Closer, error) {
	p.Lock()
	defer p.Unlock()
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrorPoolClosed
		}
		fmt.Println("Acquiring from the pool")
		return r, nil
	default:
		fmt.Println("Acquiring from the factory")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) error {
	p.Lock()
	defer p.Unlock()
	select {
	case p.resources <- resource:
		fmt.Println("Releasing the resource to the pool")
		return nil
	default:
		fmt.Println("pool full. discarding the resource")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	p.Lock()
	defer p.Unlock()
	if p.closed {
		return
	}
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
	}
}

func New(poolSize int, factory func() (io.Closer, error)) (*Pool, error) {
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolSize),
		closed:    false,
	}, nil
}
