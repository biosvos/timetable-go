package adapter

import (
	"github.com/pkg/errors"
	"timetable-go/domain"
	"timetable-go/usecase"
)

type Repository interface {
	Create(record *domain.TimeRecord) error
	Delete(id string) error
	Update(record *domain.TimeRecord) error

	List() ([]*domain.TimeRecord, error)
	Get(id string) (*domain.TimeRecord, error)
}

type SimpleUsecase struct {
	repository Repository
}

func NewSimpleUsecase(repository Repository) usecase.Usecase {
	return &SimpleUsecase{
		repository: repository,
	}
}

func (s *SimpleUsecase) CreateTimeRecord(record usecase.TimeRecord) error {
	const errMessage = "failed to create time record"

	domainRecord, err := toDomainTimeRecord(&record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	err = s.repository.Create(domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	return nil
}

func (s *SimpleUsecase) UpdateTimeRecord(record usecase.TimeRecord) error {
	const errMessage = "failed to update time record"

	domainRecord, err := toDomainTimeRecord(&record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	err = s.repository.Update(domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (s *SimpleUsecase) DeleteTimeRecord(id string) error {
	const errMessage = "failed to delete time record"

	err := s.repository.Delete(id)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (s *SimpleUsecase) GetTimeRecord(id string) (*usecase.TimeRecord, error) {
	const errMessage = "failed to get time record"

	domainRecord, err := s.repository.Get(id)
	if err != nil {
		return nil, errors.Wrap(err, errMessage)
	}

	record, err := toUsecaseRecord(domainRecord)
	if err != nil {
		return nil, errors.Wrap(err, errMessage)
	}

	return record, nil
}

func (s *SimpleUsecase) ListTimeRecords() ([]*usecase.TimeRecord, error) {
	const errMessage = "failed to list time records"

	domainRecords, err := s.repository.List()
	if err != nil {
		return nil, errors.Wrap(err, errMessage)
	}

	var ret []*usecase.TimeRecord
	for _, domainRecord := range domainRecords {
		record, err := toUsecaseRecord(domainRecord)
		if err != nil {
			return nil, errors.Wrap(err, errMessage)
		}
		ret = append(ret, record)
	}

	return ret, nil
}
