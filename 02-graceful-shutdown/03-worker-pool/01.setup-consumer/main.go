package main

import "log"

func main() {
	q := NewQueue()

	go q.startConsumer()

	for i := 0; i < 10; i++ {
		if err := q.addTask(i); err != nil {
			log.Println(err)
		}
	}
}
