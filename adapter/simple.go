package adapter

import (
	"github.com/pkg/errors"
	"timetable-go/domain"
	"timetable-go/usecase"
)

type SimpleUsecase struct {
	store map[string]*domain.TimeRecord
}

func NewSimpleUsecase() usecase.Usecase {
	return &SimpleUsecase{
		store: make(map[string]*domain.TimeRecord),
	}
}

func (s *SimpleUsecase) CreateTimeRecord(record usecase.TimeRecord) error {
	_, ok := s.store[record.Id]
	const errMessage = "failed to create time record"
	if ok {
		return errors.New(errMessage)
	}

	domainRecord, err := toDomainTimeRecord(&record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	s.store[record.Id] = domainRecord
	return nil
}

func (s *SimpleUsecase) UpdateTimeRecord(record usecase.TimeRecord) error {
	err := s.DeleteTimeRecord(record.Id)
	const errMessage = "failed to update time record"
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	err = s.CreateTimeRecord(record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (s *SimpleUsecase) DeleteTimeRecord(id string) error {
	_, ok := s.store[id]
	if !ok {
		return errors.New("failed to delete time record")
	}
	delete(s.store, id)
	return nil
}

func (s *SimpleUsecase) GetTimeRecord(id string) (*usecase.TimeRecord, error) {
	domainRecord, ok := s.store[id]
	const errMessage = "failed to get time record"
	if !ok {
		return nil, errors.New(errMessage)
	}
	record, err := toUsecaseRecord(domainRecord)
	if err != nil {
		return nil, errors.Wrap(err, errMessage)
	}

	return record, nil
}

func (s *SimpleUsecase) ListTimeRecords() ([]*usecase.TimeRecord, error) {
	var ret []*usecase.TimeRecord
	for _, domainRecord := range s.store {
		record, err := toUsecaseRecord(domainRecord)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list time records")
		}
		ret = append(ret, record)
	}

	return ret, nil
}
