package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoutesStruct struct {
	ginEngine *gin.Engine
}

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "booking service working fine!!",
		})
	})

	return r
}
