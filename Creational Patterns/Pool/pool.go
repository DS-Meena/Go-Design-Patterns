package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	Id int
}

func (w *Worker) DoWork() {
	fmt.Printf("Worker %d is doing work\n", w.Id)
}

type Pool struct {
	workers chan *Worker
	mu      sync.Mutex
	count   int
}

func NewPool(size int) *Pool {
	return &Pool{
		workers: make(chan *Worker, size),
		count:   0,
	}
}

func (p *Pool) Acquire() *Worker {
	select {
	// get a worker from channel
	case worker := <-p.workers:
		return worker

	// if no worker available, add new worker
	default:
		p.mu.Lock() // for thread safety
		defer p.mu.Unlock()
		p.count++
		return &Worker{Id: p.count}
	}
}

func (p *Pool) Release(w *Worker) {
	select {
	// add worker again to channel
	case p.workers <- w:
	default:
		fmt.Println("Pool is full, discarding worker")
	}
}
