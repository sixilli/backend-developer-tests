package concurrency

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// ErrPoolClosed is returned from AdvancedPool.Submit when the pool is closed
// before submission can be sent.
var ErrPoolClosed = errors.New("pool closed")

// AdvancedPool is a more advanced worker pool that supports cancelling the
// submission and closing the pool. All functions are safe to call from multiple
// goroutines.
type AdvancedPool interface {
	// Submit submits the given task to the pool, blocking until a slot becomes
	// available or the context is closed. The given context and its lifetime only
	// affects this function and is not the context passed to the callback. If the
	// context is closed before a slot becomes available, the context error is
	// returned. If the pool is closed before a slot becomes available,
	// ErrPoolClosed is returned. Otherwise the task is submitted to the pool and
	// no error is returned. The context passed to the callback will be closed
	// when the pool is closed.
	Submit(context.Context, func(context.Context)) error

	// Close closes the pool and waits until all submitted tasks have completed
	// before returning. If the pool is already closed, ErrPoolClosed is returned.
	// If the given context is closed before all tasks have finished, the context
	// error is returned. Otherwise, no error is returned.
	Close(context.CancelFunc) error
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.

func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	var ap AdvancedPool
	jobs := make(chan func(context.Context), maxSlots)

	wg := sync.WaitGroup{}
	wg.Add(maxConcurrent)
	sjp := SlottedJobPool{jobs, &wg, maxConcurrent, 0}

	ap = &sjp

	return ap, nil
}

type SlottedJobPool struct {
	jobs          chan func(context.Context)
	wg            *sync.WaitGroup
	maxConcurrent int
	workerCount   int
}

func (sjp *SlottedJobPool) Submit(ctx context.Context, task func(context.Context)) error {
	// Load workers if not yet initialized
	if sjp.workerCount == 0 {
		for i := 0; i < sjp.maxConcurrent; i++ {
			go AdvWorker(i, sjp.jobs, ctx, sjp.wg)
		}
		sjp.workerCount = sjp.maxConcurrent
	}

	// Submit task, if full wait.
	select {
	case <-ctx.Done():
		fmt.Println("ctx.DONE")
	case sjp.jobs <- task:
		fmt.Println("Submitted task")
	default:
		fmt.Println("Waiting to submit task")
		<-sjp.jobs
		sjp.jobs <- task
		fmt.Println("Done submit task")
	}

	return nil
}

func (sjp *SlottedJobPool) Close(closeFunc context.CancelFunc) error {
	closeFunc()
	return nil
}

func AdvWorker(id int, jobs chan func(context.Context), ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for job := range jobs {
		job := job
		fmt.Printf("worker %v starting sleep\n", id)
		job(ctx)
		fmt.Printf("worker %v awake\n", id)
	}

	return nil
}
