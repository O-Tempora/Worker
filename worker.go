package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/O-Tempora/worker/safe"
)

// Worker periodically runs given task.
//
// task - a function which worker runs over time.
//
// delay - time between finishing N-th task and running N+1-th task.
//
// onErrDelay - time between finishing N-th task and running N+1-th task if N-th task returned error.
//
// tp - provides current time for worker.
//
// ti - time interval in which worker can run its task. If nil - task can be run whenever worker is ready.
type Worker struct {
	task Task

	delay      time.Duration
	onErrDelay time.Duration

	tp TimeProvider

	ti *TimeInterval
}

// Option is a constructor option for worker.
type Option func(*Worker)

// WithDelay sets worker's delay.
func WithDelay(d time.Duration) Option {
	return func(w *Worker) { w.delay = d }
}

// WithOnErrDelay sets worker's onErrDelay.
func WithOnErrDelay(d time.Duration) Option {
	return func(w *Worker) { w.onErrDelay = d }
}

// WithCurrentTimeProvider sets worker's current time provider.
func WithCurrentTimeProvider(tp TimeProvider) Option {
	return func(w *Worker) { w.tp = tp }
}

// WithTimeInterval sets worker's task run time interval.
func WithTimeInterval(from, to Time) Option {
	return func(w *Worker) {
		ti := NewTimeInterval(from, to)
		w.ti = &ti
	}
}

// New creates new Worker.
//
// It is highly recommended to specify most of the options manualy in this constructor,
// since provided defaults may be insuffitient or inadequate in some specific usecases.
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
		tp:         defaultTimeProvider,
		ti:         nil,
	}
}

// StartBackgroundWorker starts worker's background loop of executing given task.
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
			if !w.isAllowedToRun(ctx) {
				t.Reset(w.delay)
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

func (w *Worker) isAllowedToRun(ctx context.Context) bool {
	if w.ti == nil {
		return true
	}

	return w.ti.IsInInterval(w.tp(ctx))
}
