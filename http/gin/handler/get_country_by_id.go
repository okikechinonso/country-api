package handler

import (
	"countries-api/http/gin/response"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCountryByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("id")
		log.Println(name)
		country, err := h.Db.Find(name)
		if err != nil {
			response.Failure(500, ctx, "Must provide the name of country")
			return
		}

		if country == nil {
			response.Failure(500, ctx, "Must provide the name of country")
			return
		}

		response.Success(200, *country, ctx, "Fetched Country Successfully")
	}
}
