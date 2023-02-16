package rest

import (
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"timetable-go/usecase"
)

type Rest struct {
	server  *atreugo.Atreugo
	usecase usecase.Usecase
}

func NewRest(uc usecase.Usecase) *Rest {
	config := atreugo.Config{
		Addr: "0.0.0.0:8080",
	}
	server := atreugo.New(config)
	return &Rest{
		server:  server,
		usecase: uc,
	}
}

func (r *Rest) Run() error {
	r.server.GET("/", r.CreateTimeRecord)

	err := r.server.ListenAndServe()
	return errors.Wrap(err, "failed to run")
}

func (r *Rest) CreateTimeRecord(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("hello")
}
