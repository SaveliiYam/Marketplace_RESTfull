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
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	admin := router.Group("/admin")
	{
		v1 := admin.Group("/v1")
		{
			categories := v1.Group("/categories")
			{
				categories.GET("/", h.getAllCategories)
				categories.GET("/:id", h.getCategoriesById)
				categories.POST("/", h.createCategories, h.userIdentity)
				categories.PUT("/:id", h.updateCategory, h.userIdentity)
				categories.DELETE("/:id", h.deleteCategory, h.userIdentity)
			}

			brand := v1.Group("/brands")
			{
				brand.GET("/:id", h.getById)
				brand.GET("/", h.getAllBrands)
				brand.PUT("/:id", h.updateBrand, h.userIdentity)
				brand.POST("/", h.createBrand, h.userIdentity)
				brand.DELETE("/:id", h.deleteBrand, h.userIdentity)
			}
			product := v1.Group("/products")
			{
				product.GET("/", h.getProducts)
				product.GET("/:id", h.getProductById)
				product.PUT("/:id", h.updateProduct, h.userIdentity)
				product.POST("/", h.createProduct, h.userIdentity)
				product.DELETE("/:id", h.deleteProduct, h.userIdentity)
			}
			basket := v1.Group("/basket", h.userIdentity)
			{
				basket.GET("/", h.getAllBasket)
				basket.GET("/:id", h.getBasketById)
				basket.POST("/", h.createBasket)
				basket.DELETE("/:id", h.deleteBasket)
			}
		}
	}
	return router
}
