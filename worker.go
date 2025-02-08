package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/O-Tempora/worker/safe"
)

type Worker struct {
	task Task

	delay      time.Duration
	onErrDelay time.Duration

	ti *TimeInterval
}

type Option func(*Worker)

func WithDelay(d time.Duration) Option {
	return func(w *Worker) { w.delay = d }
}

func WithOnErrDelay(d time.Duration) Option {
	return func(w *Worker) { w.onErrDelay = d }
}

func WithTimeInterval(from, to time.Time) Option {
	return func(w *Worker) {
		ti := NewTimeInterval(from, to)
		w.ti = &ti
	}
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
		task:       task,
		delay:      DefaultDelay,
		onErrDelay: DefaultOnErrDelay,
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
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if !w.isAllowedToRun() {
				continue
			}
			ctx := context.WithoutCancel(ctx)

			if err := w.runTask(ctx); err != nil {
				t.Reset(w.onErrDelay)
			}
			t.Reset(w.delay)
		}
	}
}

func (w *Worker) runTask(ctx context.Context) error {
	return w.task(ctx)
}

func (w *Worker) isAllowedToRun() bool {
	if w.ti == nil {
		return true
	}

	return w.ti.IsInInterval(time.Now())
}
