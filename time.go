package worker

import (
	"fmt"
	"time"
)

// Time represents time in hours, minutes, and seconds.
type Time struct {
	hours   int
	minutes int
	seconds int
}

func (t Time) String() string {
	return fmt.Sprintf("%d:%d:%d", t.hours, t.minutes, t.seconds)
}

func NewTime(hour, minute, second int) Time {
	return Time{
		hours:   hour,
		minutes: minute,
		seconds: second,
	}
}

func FromTime(t time.Time) Time {
	return Time{
		hours:   t.Hour(),
		minutes: t.Minute(),
		seconds: t.Second(),
	}
}

// Eq checks equality of time objects.
func (t Time) Eq(other Time) bool {
	return t == other
}

// Before returns true if t is before other.
func (t Time) Before(other Time) bool {
	if t.hours != other.hours {
		return t.hours < other.hours
	}

	if t.minutes != other.minutes {
		return t.minutes < other.minutes
	}

	return t.seconds < other.seconds
}

// After returns true if t is after other.
func (t Time) After(other Time) bool {
	if t.hours != other.hours {
		return t.hours > other.hours
	}

	if t.minutes != other.minutes {
		return t.minutes > other.minutes
	}

	return t.seconds > other.seconds
}
