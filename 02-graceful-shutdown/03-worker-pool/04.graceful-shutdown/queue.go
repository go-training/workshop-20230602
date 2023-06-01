package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	inputs    chan int
	tasks     chan int
	workerNum int
}

func (q *Queue) startConsumer(ctx context.Context) {
	for {
		select {
		case input, ok := <-q.inputs:
			if !ok {
				return
			}
			select {
			case q.tasks <- input:
			default:
				log.Println("job has been reached the limitation")
			}
		case <-ctx.Done():
			return
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

func (q *Queue) worker(ctx context.Context, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-q.tasks:
			if !ok {
				return
			}
			log.Println("worker:", num, "task:", task)
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(1000) // n will be between 0 and 10
			log.Printf("Sleeping %d Millisecond...\n", n)
			time.Sleep(time.Duration(n) * time.Millisecond)
		case <-ctx.Done():
			return
		}
	}
}

func (q *Queue) startWorker(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < q.workerNum; i++ {
		go q.worker(ctx, i, wg)
	}
}

func NewQueue() *Queue {
	return &Queue{
		inputs:    make(chan int, 100),
		tasks:     make(chan int, 1024),
		workerNum: 4,
	}
}
