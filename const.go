package worker

import (
	"context"
	"time"
)

const (
	DefaultDelay      = 500 * time.Millisecond
	DefaultOnErrDelay = 1500 * time.Second
)

var (
	defaultTimeProvider = func(_ context.Context) time.Time { return time.Now() }
)
