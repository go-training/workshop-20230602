package main

import "log"

func main() {
	q := NewQueue()

	go q.startConsumer()

	for i := 0; i < 10; i++ {
		// change the parameter for add task and print the value of variable.
		if err := q.addTask(i); err != nil {
			log.Println(err)
		}
	}
}
