package main

import (
	"log"
	"timetable-go/adapter"
	"timetable-go/infra/repository"
	"timetable-go/infra/rest"
)

func main() {
	uc := adapter.NewSimpleUsecase(repository.NewMemoryRepository())
	err := rest.NewRest(uc).Run()
	log.Println(err)
}
