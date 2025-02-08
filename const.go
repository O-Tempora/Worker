package worker

import "time"

const (
	DefaultDelay      = 500 * time.Millisecond
	DefaultOnErrDelay = 1500 * time.Second
)
