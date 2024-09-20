package main

func main() {
	pool := NewPool(2)

	for i := 0; i < 2; i++ {
		worker := pool.Acquire()
		worker.DoWork()

		worker2 := pool.Acquire()
		worker2.DoWork()

		pool.Release(worker)
		pool.Release(worker2)
	}
}

// go run *.go
// Worker 1 is doing work
// Worker 2 is doing work
// Worker 1 is doing work
// Worker 2 is doing work
