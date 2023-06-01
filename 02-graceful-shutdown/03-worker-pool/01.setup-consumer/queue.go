package main

import (
	"errors"
	"log"
)

type Queue struct {
	inputs chan int
	tasks  chan int
}

func (q *Queue) startConsumer() {
	for input := range q.inputs {
		select {
		case q.tasks <- input:
		default:
			log.Println("job has been reached the limitation")
		}
	}
}

func (q *Queue) addTask(task int) error {
	select {
	case q.inputs <- task:
	default:
		return errors.New("queue has been reached the limitation")
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{
		inputs: make(chan int, 10),
		tasks:  make(chan int, 1024),
	}
}
