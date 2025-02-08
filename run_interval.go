package worker

import "time"

type TimeInterval struct {
	from time.Time
	to   time.Time
}

func NewTimeInterval(from, to time.Time) TimeInterval {
	return TimeInterval{
		from: from,
		to:   to,
	}
}

func (tr TimeInterval) IsInInterval(t time.Time) bool {
	return t.After(tr.from) && t.Before(tr.to) ||
		t.Equal(tr.from) ||
		t.Equal(tr.to)
}
