package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	finished := make(chan struct{})
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)

	wg := &sync.WaitGroup{}
	q := NewQueue()

	go q.startConsumer(ctx)
	go q.startWorker(ctx, wg)

	wg.Add(q.workerNum)

	for i := 0; i < 10; i++ {
		if err := q.addTask(i); err != nil {
			log.Println(err)
		}
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-c:
		cancel()
		log.Println("shutdown the service:", runtime.NumGoroutine())
	}

	select {
	case <-finished:
		time.Sleep(time.Second)
		log.Println("finished the service:", runtime.NumGoroutine())
	}
}
