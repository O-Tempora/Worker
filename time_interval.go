package worker

import (
	"fmt"
	"time"
)

type TimeInterval struct {
	from Time
	to   Time
}

func (ti TimeInterval) String() string {
	return fmt.Sprintf("[%s; %s]", ti.from, ti.to)
}

func NewTimeInterval(from, to Time) TimeInterval {
	return TimeInterval{
		from: from,
		to:   to,
	}
}

func (tr TimeInterval) IsInInterval(t time.Time) bool {
	nt := FromTime(t)
	return nt.After(tr.from) && nt.Before(tr.to) ||
		nt.Eq(tr.from) ||
		nt.Eq(tr.to)
}
