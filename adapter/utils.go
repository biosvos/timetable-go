package adapter

import (
	"timetable-go/domain"
	"timetable-go/usecase"
)

func toDomainTimeRecord(record *usecase.TimeRecord) (*domain.TimeRecord, error) {
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

func toUsecaseRecord(record *domain.TimeRecord) (*usecase.TimeRecord, error) {
	return &usecase.TimeRecord{
		Id:    record.Id(),
		Start: record.StartString(),
		End:   record.EndString(),
		Memo:  record.Memo(),
	}, nil
}
