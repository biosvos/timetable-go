package rest

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/savsgio/atreugo/v11"
	"net/http"
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
	r.server.POST("/records", r.CreateTimeRecord)
	r.server.DELETE("/records/{id}", r.DeleteTimeRecord)
	r.server.PUT("/records", r.UpdateTimeRecord)
	r.server.GET("/records", r.ListTimeRecords)

	err := r.server.ListenAndServe()
	return errors.Wrap(err, "failed to run")
}

type TimeRecord struct {
	Id    string
	Start string
	End   *string
	Memo  string
}

func (r *Rest) CreateTimeRecord(ctx *atreugo.RequestCtx) error {
	var record TimeRecord
	const errMessage = "failed to create time record"
	if err := json.Unmarshal(ctx.PostBody(), &record); err != nil {
		return ctx.ErrorResponse(errors.Wrap(err, errMessage))
	}

	if err := r.usecase.CreateTimeRecord(usecase.TimeRecord{
		Id:    record.Id,
		Start: record.Start,
		End:   record.End,
		Memo:  record.Memo,
	}); err != nil {
		return ctx.ErrorResponse(errors.Wrap(err, errMessage))
	}

	return ctx.JSONResponse(nil, http.StatusNoContent)
}

func (r *Rest) DeleteTimeRecord(ctx *atreugo.RequestCtx) error {
	id := ctx.UserValue("id").(string)
	if err := r.usecase.DeleteTimeRecord(id); err != nil {
		return ctx.ErrorResponse(errors.Wrap(err, "failed to delete time record"))
	}
	return ctx.JSONResponse(nil, http.StatusOK)
}

func (r *Rest) UpdateTimeRecord(ctx *atreugo.RequestCtx) error {
	var record TimeRecord
	const errMessage = "failed to update time record"
	if err := json.Unmarshal(ctx.PostBody(), &record); err != nil {
		return ctx.ErrorResponse(errors.Wrap(err, errMessage))
	}

	if err := r.usecase.UpdateTimeRecord(usecase.TimeRecord{
		Id:    record.Id,
		Start: record.Start,
		End:   record.End,
		Memo:  record.Memo,
	}); err != nil {
		return ctx.ErrorResponse(errors.Wrap(err, errMessage))
	}

	return ctx.JSONResponse(nil, http.StatusNoContent)
}

type TimeRecordList struct {
	Items []*TimeRecord `json:"items,omitempty"`
}

func (r *Rest) ListTimeRecords(ctx *atreugo.RequestCtx) error {
	records, err := r.usecase.ListTimeRecords()
	if err != nil {
		return ctx.ErrorResponse(errors.Wrap(err, "failed to list time records"))
	}

	var items []*TimeRecord
	for _, record := range records {
		items = append(items, &TimeRecord{
			Id:    record.Id,
			Start: record.Start,
			End:   record.End,
			Memo:  record.Memo,
		})
	}
	return ctx.JSONResponse(&TimeRecordList{Items: items}, http.StatusOK)
}
