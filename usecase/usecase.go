package usecase

type TimeRecord struct {
	Id    string
	Start string
	End   *string
	Memo  string
}

type Usecase interface {
	CreateTimeRecord(record TimeRecord) error
	UpdateTimeRecord(record TimeRecord) error
	DeleteTimeRecord(id string) error

	GetTimeRecord(id string) (*TimeRecord, error)
	ListTimeRecords() ([]*TimeRecord, error)
}
