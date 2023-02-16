package domain

import (
	"github.com/pkg/errors"
	"time"
)

type Datetime struct {
	time time.Time
}

func FromString(times string) (Datetime, error) {
	parse, err := time.Parse("2006-01-02 15:04", times)
	if err == nil {
		return Datetime{time: parse}, nil
	}

	now := time.Now()
	parse, err = time.Parse("01-02 15:04", times)
	if err == nil {
		parse = parse.AddDate(now.Year(), 0, 0)
		return Datetime{time: parse}, nil
	}

	parse, err = time.Parse("02 15:04", times)
	if err == nil {
		parse = parse.AddDate(now.Year(), int(now.Month())-1, 0)
		return Datetime{time: parse}, nil
	}

	parse, err = time.Parse("15:04", times)
	if err == nil {
		parse = parse.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
		return Datetime{time: parse}, nil
	}

	parse, err = time.Parse("04", times)
	if err == nil {
		parse = parse.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
		parse = parse.Add(time.Hour * time.Duration(now.Hour()))
		return Datetime{time: parse}, nil
	}

	if len(times) == 0 {
		parse = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
		return Datetime{time: parse}, nil
	}
	const errString = "failed to from string"
	return Datetime{}, errors.Wrap(err, errString)
}

func (d Datetime) String() string {
	return d.time.Format("2006-01-02 15:04")
}
