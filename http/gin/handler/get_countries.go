package handler

import (
	"countries-api/http/gin/response"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCountries() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		countries, err := h.Db.FindMany()
		if err != nil {
			log.Println(err)
			response.Failure(400, ctx, "enter a valid number")
			return
		}

		response.Success(200, countries, ctx, "Fetched countries successfull")
	}
}
