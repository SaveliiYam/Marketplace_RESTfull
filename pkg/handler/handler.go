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
			categories := v1.Group("/categories")
			{
				categories.GET("/", h.getCategories)
				categories.GET("/:id", h.getCategoriesById)
				categories.POST("/", h.createCategories)
				categories.PUT("/:id", h.updateCategory)
				categories.DELETE("/:id", h.deleteCategory)
			}

			brand := v1.Group("/brands")
			{
				brand.GET("/:id", h.getById)
				brand.GET("/", h.getAllBrands)
				brand.PUT("/:id", h.updateBrand)
				brand.POST("/", h.createBrand)
				brand.DELETE("/:id", h.deleteBrand)
			}
			product := v1.Group("/products")
			{
				product.GET("/", h.getProducts)
				product.GET("/:id", h.getProductById)
				product.PUT("/:id", h.updateProduct)
				product.POST("/", h.createProduct)
				product.DELETE("/:id", h.deleteProduct)
			}
		}
	}
	return router
}
