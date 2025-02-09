package worker_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/O-Tempora/worker"
)

func taskCounter(count uint32) (worker.Task, <-chan struct{}) {
	var ct atomic.Uint32
	doneCh := make(chan struct{})

	task := worker.Task(func(ctx context.Context) error {
		if ct.Add(1) == count {
			close(doneCh)
		}

		return nil
	})

	return task, doneCh
}

func TestWorker_Basics(t *testing.T) {
	t.Parallel()

	t.Run("expecting to run at least 5 times in a second with no limitations", func(t *testing.T) {
		t.Parallel()

		tsk, done := taskCounter(5)
		wk := worker.New(tsk, worker.WithDelay(0))

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err := worker.StartBackgroundWorker(ctx, wk)
		if err != nil {
			t.Fatalf("worker error: %s", err.Error())
		}

		select {
		case <-ctx.Done():
			t.Fatal("context Done() had fired before task execution count was reached")
		case <-done:
			// noop
		}
	})

	t.Run("expecting to stop after parent context cancelation", func(t *testing.T) {
		t.Parallel()

		tsk, done := taskCounter(0)
		wk := worker.New(tsk, worker.WithDelay(0))

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		cancel()

		err := worker.StartBackgroundWorker(ctx, wk)
		if err != nil {
			t.Fatalf("worker error: %s", err.Error())
		}

		select {
		case <-ctx.Done():
			// noop
		case <-done:
			t.Fatal("parent context is canceled therefore task must not be executed")
		}
	})

	t.Run("expecting to run exactly once in a second with 0.8 sec delay", func(t *testing.T) {
		t.Parallel()

		tsk, done := taskCounter(2)
		wk := worker.New(tsk, worker.WithDelay(1*time.Second))

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		cancel()

		err := worker.StartBackgroundWorker(ctx, wk)
		if err != nil {
			t.Fatalf("worker error: %s", err.Error())
		}

		select {
		case <-ctx.Done():
			// noop
		case <-done:
			t.Fatal("context must be cancelled because of delay")
		}
	})
}

func TestWorker_Validation(t *testing.T) {
	t.Parallel()

	t.Run("worker is nil", func(t *testing.T) {
		t.Parallel()

		err := worker.StartBackgroundWorker(context.Background(), nil)
		if err == nil {
			t.Fatal("must return an error since worker is nil")
		}
	})

	t.Run("task is nil", func(t *testing.T) {
		t.Parallel()

		err := worker.StartBackgroundWorker(
			context.Background(),
			worker.New(nil),
		)
		if err == nil {
			t.Fatal("must return an error since task is nil")
		}
	})
}

func TestWorker_IsNotInRunInterval(t *testing.T) {
	t.Parallel()

	tsk, done := taskCounter(1)
	wk := worker.New(
		tsk,
		worker.WithDelay(1*time.Second),
		worker.WithCurrentTimeProvider(func(ctx context.Context) time.Time {
			return time.Date(1990, 5, 5, 10, 10, 10, 1, time.UTC)
		}),
		worker.WithTimeInterval(
			worker.NewTime(22, 0, 0),
			worker.NewTime(23, 0, 0),
		),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	cancel()

	err := worker.StartBackgroundWorker(ctx, wk)
	if err != nil {
		t.Fatalf("worker error: %s", err.Error())
	}

	select {
	case <-ctx.Done():
		// noop
	case <-done:
		t.Fatal("context must be cancelled because of delay")
	}
}
