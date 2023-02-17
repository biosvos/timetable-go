package main

import (
	"log"
	"timetable-go/adapter"
	"timetable-go/infra/broker"
	"timetable-go/infra/repository"
	"timetable-go/infra/rest"
)

func main() {
	fileRepository, err := repository.NewFileRepository(repository.NewMemoryRepository(), "testfile")
	if err != nil {
		log.Fatal(err)
	}

	zeromqBroker, closeFn, err := broker.NewZeroMessageQueueBroker(5555)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFn()

	uc := adapter.NewSimpleUsecase(fileRepository, zeromqBroker)
	err = rest.NewRest(uc).Run()
	if err != nil {
		log.Fatal(err)
	}
}
