package safe

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
)

// Call runs input function fn and recovers a panic (if any happened).
func Call(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered while trying to do safe.Call, stack: %s", string(debug.Stack()))
			err = fmt.Errorf("panic recovered: %+v", r)
		}
	}()
	return fn(ctx)
}

// Go is the asynchronous version of Call.
// It does not perform error check.
func Go(ctx context.Context, fn func(ctx context.Context)) {
	go func() {
		_ = Call(ctx, func(ctx context.Context) error {
			fn(ctx)
			return nil
		})
	}()
}
