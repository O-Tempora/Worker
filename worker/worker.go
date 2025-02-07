package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/O-Tempora/hive/safe"
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

func StartBackgroundWorker(ctx context.Context, w *Worker) error {
	if err := w.validate(); err != nil {
		return fmt.Errorf("worker validate: %w", err)
	}

	safe.Go(ctx, w.run)

	return nil
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

	t := time.NewTimer(0)
	defer t.Stop()

	for {
		time.Sleep(w.delay)
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			ctx := context.Background()
			if err := w.runTask(ctx); err != nil {
				log.Println("finished work with error: ", err.Error())
			}
			t.Reset(w.delay)
			continue
		}
	}
}

func (w *Worker) runTask(ctx context.Context) error {
	startedAt := time.Now()
	log.Printf("started task at %s\n", startedAt.Format(time.RFC3339))

	return w.task(ctx)
}
