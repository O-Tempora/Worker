package worker

import (
	"fmt"
	"time"
)

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

func (t Time) Eq(other Time) bool {
	return t == other
}

func (t Time) Before(other Time) bool {
	if t.hours != other.hours {
		return t.hours < other.hours
	}

	if t.minutes != other.minutes {
		return t.minutes < other.minutes
	}

	return t.seconds < other.seconds
}

func (t Time) After(other Time) bool {
	if t.hours != other.hours {
		return t.hours > other.hours
	}

	if t.minutes != other.minutes {
		return t.minutes > other.minutes
	}

	return t.seconds > other.seconds
}
