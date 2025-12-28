package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piyushsharma67/events_booking/services/events_service/middleware"
	"github.com/piyushsharma67/events_booking/services/events_service/service"
)

type RoutesStruct struct {
	ginEngine *gin.Engine
	service   *service.EventService
}

func (r *RoutesStruct) InitialiseRoutes() *gin.Engine {

	if r.ginEngine != nil {
		r.ginEngine = gin.Default()

		return r.ginEngine
	}

	return r.ginEngine

}

func InitRoutes(service *service.EventService) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "I am working fine!!",
		})
	})

	organise := r.Group("organize")
	organise.Use(middleware.RoleAuthMiddleware("organizer"))

	organise.POST("/organise/create")

	return r
}
