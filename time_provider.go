package worker

import (
	"context"
	"time"
)

// TimeProvider is a provider of current time for worker.
type TimeProvider func(ctx context.Context) time.Time
