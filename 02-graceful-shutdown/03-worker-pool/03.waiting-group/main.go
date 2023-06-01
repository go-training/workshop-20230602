package main

import (
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	q := NewQueue()

	go q.startConsumer()
	go q.startWorker(wg)

	wg.Add(q.workerNum)

	for i := 0; i < 100; i++ {
		if err := q.addTask(i); err != nil {
			log.Println(err)
		}
	}

	wg.Wait()
}
