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

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				categories := user.Group("/categories")
				{
					categories.GET("/", h.getAllCategories)
					categories.GET("/:id", h.getCategoriesById)
					categories.GET("/:id/image", h.getCategoriesImage)
				}

				brand := user.Group("/brands")
				{
					brand.GET("/:id", h.getById)
					brand.GET("/", h.getAllBrands)
					brand.GET("/:id/image", h.getBrandsImage)
				}
				product := user.Group("/products")
				{
					product.GET("/", h.getProducts)
					product.GET("/:id", h.getProductById)
					product.GET(":/id/image", h.getProductImage)
				}
				basket := user.Group("/basket", h.userIdentity)
				{
					basket.GET("/", h.getAllBasket)
					basket.GET("/:id", h.getBasketById)
					basket.POST("/", h.createBasket)
					basket.DELETE("/:id", h.deleteBasket)
				}
			}
			admin := v1.Group("/admin")
			{
				categories := admin.Group("/categories", h.userIdentity)
				{
					categories.GET("/", h.getAllCategories)
					categories.GET("/:id", h.getCategoriesById)
					categories.GET("/:id/image", h.getCategoriesImage)
					categories.POST("/", h.createCategories)
					categories.POST("/:id/upload", h.createCategoriesImage)
					categories.PUT("/:id", h.updateCategory)
					categories.DELETE("/:id", h.deleteCategory)
				}

				brand := admin.Group("/brands", h.userIdentity)
				{
					brand.GET("/:id", h.getById)
					brand.GET("/", h.getAllBrands)
					brand.GET("/:id/image", h.getBrandsImage)
					brand.PUT("/:id", h.updateBrand)
					brand.POST("/", h.createBrand)
					brand.POST("/:id/upload", h.createBrandsImage)
					brand.DELETE("/:id", h.deleteBrand)
				}
				product := admin.Group("/products", h.userIdentity)
				{
					product.GET("/", h.getProducts)
					product.GET("/:id", h.getProductById)
					product.GET(":/id/image", h.getProductImage)
					product.PUT("/:id", h.updateProduct)
					product.POST("/", h.createProduct)
					product.POST(":/id/image", h.createProductImage)
					product.DELETE("/:id", h.deleteProduct)
				}
				basket := admin.Group("/basket", h.userIdentity)
				{
					basket.GET("/", h.getAllBasket)
					basket.GET("/:id", h.getBasketById)
					basket.POST("/", h.createBasket)
					basket.DELETE("/:id", h.deleteBasket)
				}
			}
		}
	}

	return router
}
