package worker

import "context"

// Task is a function, that is expected to be executed by worker.
type Task func(ctx context.Context) error
