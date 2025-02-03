package worker

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Worker struct {
	task Task
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

func StartBackgroundWorker(ctx context.Context, w *Worker) error {
	if err := w.validate(); err != nil {
		return fmt.Errorf("worker validate: %w", err)
	}

	// set defaults
	go func() {
		w.run(ctx)
	}()

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
			continue
		}
	}
}

func (w *Worker) runTask(ctx context.Context) error {
	startedAr := time.Now()
	log.Printf("started task at %s\n", startedAr.Format(time.RFC3339))

	return w.task(ctx)
}
