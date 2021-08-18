package handler

import (
	"dogma_test_task/internal/service"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "dogma_test_task/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()

	app.Use(gin.Logger())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := app.Group("/api")
	{
		lists := api.Group("/users")
		{
			lists.GET("/", h.GetUserList)
			lists.POST("/", h.AddUser)
			lists.GET("/:id", h.GetUserById)
			lists.PUT("/:id", h.UpdateUser)
			lists.DELETE("/:id", h.DeleteUser)

		}
	}

	return app
}
