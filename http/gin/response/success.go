package response

import "github.com/gin-gonic/gin"

func Success(code int, data interface{}, c *gin.Context, message string) {
	c.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}
