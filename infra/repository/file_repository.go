package repository

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"timetable-go/adapter"
	"timetable-go/domain"
)

type row struct {
	Id    string
	Start string
	End   *string
	Memo  string
}

type fileRepository struct {
	other    adapter.Repository
	filename string
}

func NewFileRepository(repository adapter.Repository, filename string) (adapter.Repository, error) {
	ret := fileRepository{
		other:    repository,
		filename: filename,
	}
	err := ret.load()
	if err != nil {
		return nil, errors.Wrap(err, "failed to new file repository")
	}
	return &ret, nil
}

func toRow(record *domain.TimeRecord) *row {
	return &row{
		Id:    record.Id(),
		Start: record.StartString(),
		End:   record.EndString(),
		Memo:  record.Memo(),
	}
}

func (f *fileRepository) save() error {
	const errMessage = "failed to save records"

	list, err := f.other.List()
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	var ret []*row
	for _, record := range list {
		rows := toRow(record)
		ret = append(ret, rows)
	}

	marshal, err := json.Marshal(ret)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	err = os.WriteFile(f.filename, marshal, 0)
	return errors.Wrap(err, errMessage)
}

func toDomainTimeRecord(record *row) (*domain.TimeRecord, error) {
	start, err := domain.FromString(record.Start)
	if err != nil {
		return nil, err
	}
	domainRecord := domain.NewTimeRecord(record.Id, start, record.Memo)
	if record.End == nil {
		return domainRecord, nil
	}

	end, err := domain.FromString(*record.End)
	if err != nil {
		return nil, err
	}
	domainRecord, err = domainRecord.WithEnd(end)
	if err != nil {
		return nil, err
	}
	return domainRecord, nil
}

func (f *fileRepository) load() error {
	const errMessage = "failed to load records"

	bytes, err := os.ReadFile(f.filename)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, errMessage)
	}

	var rows []*row
	err = json.Unmarshal(bytes, &rows)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}

	for _, item := range rows {
		record, err := toDomainTimeRecord(item)
		if err != nil {
			return errors.Wrap(err, errMessage)
		}
		err = f.Create(record)
		if err != nil {
			return errors.Wrap(err, errMessage)
		}
	}
	return nil
}

func (f *fileRepository) Create(record *domain.TimeRecord) error {
	const errMessage = "failed to create record"
	err := f.other.Create(record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	if err := f.save(); err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (f *fileRepository) Delete(id string) error {
	const errMessage = "failed to delete record"
	err := f.other.Delete(id)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	if err := f.save(); err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (f *fileRepository) Update(record *domain.TimeRecord) error {
	const errMessage = "failed to update record"
	err := f.other.Update(record)
	if err != nil {
		return errors.Wrap(err, errMessage)
	}
	if err := f.save(); err != nil {
		return errors.Wrap(err, errMessage)
	}
	return nil
}

func (f *fileRepository) List() ([]*domain.TimeRecord, error) {
	const errMessage = "failed to list records"
	list, err := f.other.List()
	if err != nil {
		return nil, errors.Wrap(err, errMessage)
	}
	return list, nil
}

func (f *fileRepository) Get(id string) (*domain.TimeRecord, error) {
	const errMessage = "failed to get record"
	record, err := f.other.Get(id)
	if err != nil {
		return nil, errors.Wrap(err, errMessage)
	}
	return record, nil
}
