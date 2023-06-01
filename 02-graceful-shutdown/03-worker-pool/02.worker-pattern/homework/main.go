package main

import (
	"log"
	"time"
)

func main() {
	// update default worker number
	// update input chan size
	// update task chan size
	q := NewQueue()

	go q.startConsumer()
	go q.startWorker()

	for i := 0; i < 100; i++ {
		if err := q.addTask(i); err != nil {
			log.Println(err)
		}
	}

	time.Sleep(2 * time.Second)
}
