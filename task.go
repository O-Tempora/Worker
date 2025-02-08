package worker

import "context"

type Task func(ctx context.Context) error
