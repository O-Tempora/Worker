package safe

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
)

func Call(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered while trying to do safe.Call, stack: %s", string(debug.Stack()))
			err = fmt.Errorf("panic recovered: %+v", r)
		}
	}()
	return fn(ctx)
}

func Go(ctx context.Context, fn func(ctx context.Context)) {
	go func() {
		Call(ctx, func(ctx context.Context) error {
			fn(ctx)
			return nil
		})
	}()
}
