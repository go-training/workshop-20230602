package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type option func(*Queue)

type Queue struct {
	inputs    chan int
	tasks     chan int
	workerNum int
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

func (q *Queue) worker(num int) {
	for task := range q.tasks {
		log.Println("worker:", num, "task:", task)
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000) // n will be between 0 and 10
		log.Printf("Sleeping %d Millisecond...\n", n)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func (q *Queue) startWorker() {
	for i := 0; i < q.workerNum; i++ {
		go q.worker(i)
	}
}

// support custom parameter.
func NewQueue(opts ...option) *Queue {
	return &Queue{
		inputs:    make(chan int, 100),
		tasks:     make(chan int, 1024),
		workerNum: 4,
	}
}
