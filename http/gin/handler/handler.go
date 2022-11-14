package handler

import (
	"countries-api/db"
	"countries-api/domain"
)

type Handler struct {
	Db domain.DbInterface
}

func NewHandler() *Handler {
	return &Handler{Db: db.NewDb()}
}
