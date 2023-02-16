package repository

import (
	"github.com/pkg/errors"
	"timetable-go/adapter"
	"timetable-go/domain"
)

type memoryRepository struct {
	store map[string]*domain.TimeRecord
}

func NewMemoryRepository() adapter.Repository {
	return &memoryRepository{
		store: make(map[string]*domain.TimeRecord),
	}
}

func (m *memoryRepository) Create(record *domain.TimeRecord) error {
	_, ok := m.store[record.Id()]
	const errMessage = "failed to create record"
	if ok {
		return errors.New(errMessage)
	}
	m.store[record.Id()] = record
	return nil
}

func (m *memoryRepository) Delete(id string) error {
	_, ok := m.store[id]
	const errMessage = "failed to delete time record"
	if !ok {
		return errors.New(errMessage)
	}
	delete(m.store, id)
	return nil
}

func (m *memoryRepository) Update(record *domain.TimeRecord) error {
	_, ok := m.store[record.Id()]
	const errMessage = "failed to update time record"
	if !ok {
		return errors.New(errMessage)
	}
	m.store[record.Id()] = record
	return nil
}

func (m *memoryRepository) List() ([]*domain.TimeRecord, error) {
	var ret []*domain.TimeRecord
	for _, record := range m.store {
		ret = append(ret, record)
	}
	return ret, nil
}

func (m *memoryRepository) Get(id string) (*domain.TimeRecord, error) {
	record, ok := m.store[id]
	const errMessage = "failed to get record"
	if !ok {
		return nil, errors.New(errMessage)
	}
	return record, nil
}
