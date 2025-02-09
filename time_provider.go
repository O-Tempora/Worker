package worker

import (
	"context"
	"time"
)

type TimeProvider func(ctx context.Context) time.Time
