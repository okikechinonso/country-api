package entity

import (
	"time"

	"github.com/google/uuid"
)

type Country struct {
	ID            string    `json:"id" bson:"_id"`
	Name          string    `json:"name" bson:"name" binding:"required"`
	ShortName     string    `json:"short_name" bson:"short_name" binding:"required"`
	Continent     string    `json:"continent" bson:"continent" binding:"required"`
	IsOperational bool      `json:"is_operational" bson:"is_operational" binding:"required"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at" `
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}

func NewCountry() *Country {
	return &Country{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
