package worker_test

import (
	"testing"

	"github.com/O-Tempora/worker"
)

func TestTime_Comparison(t *testing.T) {
	t.Parallel()

	t.Run("equal", func(t *testing.T) {
		t.Parallel()

		t1, t2 := worker.NewTime(1, 2, 3), worker.NewTime(1, 2, 3)
		if !t1.Eq(t2) {
			t.Errorf("%s must be equal to %s", t1, t2)
		}
		if t1.Before(t2) {
			t.Errorf("%s must not be before to %s", t1, t2)
		}
		if t1.After(t2) {
			t.Errorf("%s must not be after to %s", t1, t2)
		}
	})

	t.Run("t1 is before t2", func(t *testing.T) {
		t.Parallel()

		t1, t2 := worker.NewTime(1, 2, 3), worker.NewTime(2, 0, 0)
		if !t1.Before(t2) {
			t.Errorf("%s must be before to %s", t1, t2)
		}
	})

	t.Run("t1 is after t2", func(t *testing.T) {
		t.Parallel()

		t1, t2 := worker.NewTime(2, 1, 0), worker.NewTime(2, 0, 0)
		if !t1.After(t2) {
			t.Errorf("%s must be after to %s", t1, t2)
		}
	})
}
