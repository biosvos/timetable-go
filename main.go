package main

import (
	"log"
	"timetable-go/adapter"
	"timetable-go/infra/rest"
)

func main() {
	uc := adapter.NewSimpleUsecase()
	err := rest.NewRest(uc).Run()
	log.Println(err)
}
