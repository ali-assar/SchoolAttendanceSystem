package handler

import "github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"

type Handlers struct {
	Store db.Querier
}

func NewHandlers(store db.Querier) *Handlers {
	return &Handlers{
		Store: store,
	}
}
