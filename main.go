package main

import (
	"log"
	"timetable-go/adapter"
	"timetable-go/usecase"
)

func main() {
	uc := adapter.NewSimpleUsecase()
	err := uc.CreateTimeRecord(usecase.TimeRecord{
		Id:    "asdf",
		Start: "10",
		End:   nil,
		Memo:  "joidsf",
	})
	log.Println(err)
	records, err := uc.ListTimeRecords()
	log.Println(err)
	for _, record := range records {
		log.Println(record.Start, "~", record.End, ":", record.Memo)
	}
}
