package handler

import (
	"marketplace/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signIn)
		auth.POST("sign-in", h.signUp)
	}
	admin := router.Group("/admin")
	{
		v1 := admin.Group("/v1")
		{
			v1.GET("/categories", h.getCategories)

			shoes := v1.Group("/shoes")
			{
				shoes.PUT("/:id")
				shoes.POST("/")
			}
			cloth := v1.Group("/shoes")
			{
				cloth.PUT("/:id")
				cloth.POST("/")
			}
			accessories := v1.Group("/accessories")
			{
				accessories.PUT("/:id")
				accessories.POST("/")
			}
			brand := v1.Group("/brands")
			{
				brand.PUT("/:id")
				brand.POST("/")
			}
		}
	}
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/categories", h.getCategories)

			shoes := v1.Group("/shoes")
			{
				shoes.GET("/:id")
				shoes.GET("/")
			}
			cloth := v1.Group("/cloth")
			{
				cloth.GET("/:id")
				cloth.GET("/")
			}
			accessories := v1.Group("/accessories")
			{
				accessories.GET("/:id")
				accessories.GET("/")
			}
			brand := v1.Group("/brands")
			{
				brand.GET("/:id")
				brand.GET("/")
			}
		}
	}

	return router
}
