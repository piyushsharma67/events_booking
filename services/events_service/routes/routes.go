package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoutesStruct struct{
	ginEngine *gin.Engine
}

func (r *RoutesStruct)InitialiseRoutes()*gin.Engine{

	if r.ginEngine!=nil{
		r.ginEngine=gin.Default()

		return r.ginEngine
	}

	return r.ginEngine
	
}

func InitRoutes() *gin.Engine {
	r:=gin.Default()
	r.Use(gin.Logger())
	r.GET("/health", func(ctx *gin.Context) {
		fmt.Println("i am called")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "I am working fine!!",
		})
	})

	return r
}