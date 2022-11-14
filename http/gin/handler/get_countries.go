package handler

import (
	"countries-api/http/gin/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (h *Handler) GetCountries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Query("page"))
		log.Println(page)
		if err != nil {
			response.Failure(400, ctx, "enter a valid number")
			return
		}

		countries, err := h.Db.FindMany(page)
		if err != nil {
			log.Println(err)
			response.Failure(400, ctx, "enter a valid number")
			return
		}

		response.Success(200, countries, ctx, "Fetched countries successfull")
	}
}
