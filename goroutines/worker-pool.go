package goroutine

import (
	"context"
	"fmt"
	"time"
)

type WorkerPoolInterface interface {
	ConsumeEvent(ctx context.Context, evt *Event) error
	// NewWorkerPool(ctx context.Context, workerCount int, timeout time.Time) (*WorkerPool, error)
}

type Event struct {
	msg string
}

type WorkerPool struct {
	WorkerPoolInterface
	maxWorkerCount    int
	timeout           time.Duration
	remainingWorkers  int
	workerDoneChannel chan int
}

func NewWorkerPool(ctx context.Context, workerCount int, timeout time.Duration) (*WorkerPool, error) {
	ch := make(chan int, workerCount)
	wp := &WorkerPool{
		maxWorkerCount:    workerCount,
		timeout:           timeout,
		remainingWorkers:  workerCount,
		workerDoneChannel: ch,
	}
	return wp, nil
}

func (w *WorkerPool) ConsumeEvent(ctx context.Context, evt *Event) error {
	if w.remainingWorkers > w.maxWorkerCount {
		w.remainingWorkers = w.maxWorkerCount
	}
remainingWorkerPoint:
	if w.remainingWorkers > 0 {
		w.remainingWorkers -= 1
		go func() {
			handleEvent(ctx, evt)
			w.remainingWorkers += 1
			w.workerDoneChannel <- 1
		}()
	} else {
		select {
		case <-w.workerDoneChannel:
			goto remainingWorkerPoint
		}
	}
	return nil
}

func handleEvent(ctx context.Context, evt *Event) {

	fmt.Println(evt.msg, time.Now())
	time.Sleep(1)
}


