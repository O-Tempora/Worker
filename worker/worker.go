package worker

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Worker struct {
	task  Task
	delay time.Duration
}

type Option func(*Worker)

func WithDelay(d time.Duration) Option {
	return func(w *Worker) { w.delay = d }
}

func New(task Task, opts ...Option) *Worker {
	w := newWorker(task)
	for i := range opts {
		opts[i](w)
	}
	return w
}

func newWorker(task Task) *Worker {
	return &Worker{
		delay: DefaultDelay,
		task:  task,
	}
}

func (w *Worker) validate() error {
	if w == nil {
		return fmt.Errorf("worker is nil")
	}

	if w.task == nil {
		return fmt.Errorf("task is nil")
	}

	return nil
}

func (w *Worker) run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ctx := context.Background()
			if err := w.runTask(ctx); err != nil {
				log.Println("finished work with error: ", err.Error())
			}

			time.Sleep(w.delay)
			continue
		}
	}
}

func (w *Worker) runTask(ctx context.Context) error {
	startedAt := time.Now()
	log.Printf("started task at %s\n", startedAt.Format(time.RFC3339))

	return w.task(ctx)
}

func StartBackgroundWorker(ctx context.Context, w *Worker) error {
	if err := w.validate(); err != nil {
		return fmt.Errorf("worker validate: %w", err)
	}

	go w.run(ctx)

	return nil
}
