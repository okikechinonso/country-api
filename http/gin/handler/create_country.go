package handler

import (
	"countries-api/entity"
	"countries-api/http/gin/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCountry() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		country := entity.NewCountry()
		err := ctx.BindJSON(country)
		country.Name = strings.ToLower(country.Name)
		if err != nil {
			response.Failure(500, ctx, "unable to decode data")
			return
		}

		_, err = h.Db.Create(*country)
		if err != nil {
			response.Failure(500, ctx, "unable to create")
			return
		}

		response.Success(200, country, ctx, "Country created successful")
	}
}
