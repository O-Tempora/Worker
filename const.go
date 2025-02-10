package worker

import (
	"context"
	"time"
)

const (
	// DefaultDelay is a default delay between task runs.
	DefaultDelay = 500 * time.Millisecond
	// DefaultOnErrDelay is a default delay between tasks if an error occurs.
	DefaultOnErrDelay = 3 * time.Second
)

var (
	defaultTimeProvider = func(_ context.Context) time.Time { return time.Now() }
)
