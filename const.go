package worker

import (
	"context"
	"time"
)

const (
	// DefaultDelay is a default delay between task runs.
	DefaultDelay = 500 * time.Millisecond
	// DefaultOnErrDelay is a default delay between tasks if an error occurs.
	DefaultOnErrDelay = 2 * time.Second
	// DefaultRunTimeout is a default timeout for worker's task.
	DefaultRunTimeout = 3 * time.Second
)

var (
	// DefaultTimeProvider is a time provider that always returns time.Now().
	DefaultTimeProvider = func(_ context.Context) time.Time { return time.Now() }
)
