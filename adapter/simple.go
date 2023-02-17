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

type Topic string

const (
	TopicCreated = Topic("created")
	TopicDeleted = Topic("deleted")
	TopicUpdated = Topic("updated")
)

type Broker interface {
	Publish(topic Topic, record *domain.TimeRecord) error
}

type simpleUsecase struct {
	repository Repository
	broker     Broker
}

func NewSimpleUsecase(repository Repository, broker Broker) usecase.Usecase {
	return &simpleUsecase{
		repository: repository,
		broker:     broker,
	}
}

func (s *simpleUsecase) CreateTimeRecord(record usecase.TimeRecord) error {
	const errMessage = "failed to create time record"

	domainRecord, err := toDomainTimeRecord(&record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	err = s.repository.Create(domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	err = s.broker.Publish(TopicCreated, domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (s *simpleUsecase) UpdateTimeRecord(record usecase.TimeRecord) error {
	const errMessage = "failed to update time record"

	domainRecord, err := toDomainTimeRecord(&record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	err = s.repository.Update(domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	err = s.broker.Publish(TopicUpdated, domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (s *simpleUsecase) DeleteTimeRecord(id string) error {
	const errMessage = "failed to delete time record"
	domainRecord, err := s.repository.Get(id)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	err = s.repository.Delete(id)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	err = s.broker.Publish(TopicDeleted, domainRecord)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (s *simpleUsecase) GetTimeRecord(id string) (*usecase.TimeRecord, error) {
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

func (s *simpleUsecase) ListTimeRecords() ([]*usecase.TimeRecord, error) {
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
