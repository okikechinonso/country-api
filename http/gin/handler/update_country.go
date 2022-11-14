package handler

import (
	"countries-api/http/gin/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Dto struct {
	Name          string `json:"name" bson:"name" binding:"required"`
	ShortName     string `json:"short_name" bson:"short_name" binding:"required"`
	Continent     string `json:"continent" bson:"continent" binding:"required"`
	IsOperational bool   `json:"is_operational" bson:"is_operational"`
}

func (h *Handler) UpdateCountry() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("id")
		body := &Dto{}
		err := ctx.ShouldBindJSON(body)
		if err != nil {
			log.Println(err)
			response.Failure(400, ctx, "provide all required fields")
			return
		}
		country, err := h.Db.Find(name)
		if err != nil {
			response.Failure(400, ctx, "user doesn't exist")
			return
		}

		country.Name, country.Continent, country.ShortName, country.IsOperational = body.Name, body.Continent, body.ShortName, body.IsOperational

		data, err := h.Db.Update(*country, country.ID)
		log.Print(*country)
		if err != nil {
			response.Failure(500, ctx, "Error occurred when processing your request")
			return
		}

		response.Success(200, data, ctx, fmt.Sprintf("update %v successful", country.ID))
	}
}
