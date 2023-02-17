package main

import (
	"log"
	"timetable-go/adapter"
	"timetable-go/infra/repository"
	"timetable-go/infra/rest"
)

func main() {
	fileRepository, err := repository.NewFileRepository(repository.NewMemoryRepository(), "testfile")
	if err != nil {
		log.Fatal(err)
	}
	uc := adapter.NewSimpleUsecase(fileRepository)
	err = rest.NewRest(uc).Run()
	if err != nil {
		log.Fatal(err)
	}
}
