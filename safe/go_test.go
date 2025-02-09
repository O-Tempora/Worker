package safe_test

import (
	"context"
	"testing"

	"github.com/O-Tempora/worker/safe"
)

func TestSafeCall_NoPanic(t *testing.T) {
	t.Parallel()

	fn := func(ctx context.Context) error {
		panic("some panic")
	}

	if err := safe.Call(context.Background(), fn); err == nil {
		t.Error("expected error, got nil")
	}
}
