package domain

import "github.com/pkg/errors"

type TimeRecord struct {
	id    string
	start Datetime
	end   *Datetime
	memo  string
}

func (t *TimeRecord) Id() string {
	return t.id
}

func (t *TimeRecord) Start() Datetime {
	return t.start
}

func (t *TimeRecord) StartString() string {
	return t.start.String()
}

func (t *TimeRecord) End() *Datetime {
	return t.end
}

func (t *TimeRecord) EndString() *string {
	if t.end == nil {
		return nil
	}
	ret := t.end.String()
	return &ret
}

func (t *TimeRecord) Memo() string {
	return t.memo
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
