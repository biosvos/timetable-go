package domain

import "github.com/pkg/errors"

type TimeRecord struct {
	id    string
	start Datetime
	end   *Datetime
	memo  string
}

func NewTimeRecord(id string, start Datetime, memo string) *TimeRecord {
	return &TimeRecord{
		id:    id,
		start: start,
		memo:  memo,
	}
}

func (t *TimeRecord) WithEnd(end Datetime) (*TimeRecord, error) {
	if t.start.time.After(end.time) {
		return nil, errors.New("must start time <= end time")
	}
	ret := *t
	ret.end = &end
	return &ret, nil
}
