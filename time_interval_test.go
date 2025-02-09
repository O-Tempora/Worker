package worker_test

import (
	"testing"
	"time"

	"github.com/O-Tempora/worker"
)

func TestTimeInterval_IsInInterval(t *testing.T) {
	t.Parallel()

	t.Run("is in interval", func(t *testing.T) {
		t.Parallel()

		tmstmp := time.Date(1990, 5, 5, 10, 10, 10, 1, time.UTC)
		lb, rb := worker.FromTime(tmstmp.Add(-1*time.Hour)), worker.FromTime(tmstmp.Add(1*time.Hour))

		if interval := worker.NewTimeInterval(lb, rb); !interval.IsInInterval(tmstmp) {
			t.Errorf("expected time %s to be inside interval %s", tmstmp.Format(time.TimeOnly), interval)
		}
	})

	t.Run("is not in interval", func(t *testing.T) {
		t.Parallel()

		tmstmp := time.Date(1990, 5, 5, 10, 10, 10, 1, time.UTC)
		lb, rb := worker.FromTime(tmstmp.Add(2*time.Hour)), worker.FromTime(tmstmp.Add(3*time.Hour))

		if interval := worker.NewTimeInterval(lb, rb); interval.IsInInterval(tmstmp) {
			t.Errorf("expected time %s to be outside of interval %s", tmstmp.Format(time.TimeOnly), interval)
		}
	})

	t.Run("is on left border", func(t *testing.T) {
		t.Parallel()

		tmstmp := time.Date(1990, 5, 5, 10, 10, 10, 1, time.UTC)
		lb, rb := worker.FromTime(tmstmp), worker.FromTime(tmstmp.Add(1*time.Hour))

		if interval := worker.NewTimeInterval(lb, rb); !interval.IsInInterval(tmstmp) {
			t.Errorf("expected time %s to be inside interval %s (on left border)", tmstmp.Format(time.TimeOnly), interval)
		}
	})

	t.Run("is on right border", func(t *testing.T) {
		t.Parallel()

		tmstmp := time.Date(1990, 5, 5, 10, 10, 10, 1, time.UTC)
		lb, rb := worker.FromTime(tmstmp.Add(-2*time.Hour)), worker.FromTime(tmstmp)

		if interval := worker.NewTimeInterval(lb, rb); !interval.IsInInterval(tmstmp) {
			t.Errorf("expected time %s to be inside interval %s (on right border)", tmstmp.Format(time.TimeOnly), interval)
		}
	})
}
