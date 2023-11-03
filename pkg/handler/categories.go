package handler

import (
	"marketplace"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllCategoriesData struct {
	Data []marketplace.CategoriesList `json:"data"`
}

func (h *Handler) getCategories(c *gin.Context) {
	categories, err := h.services.Categories.GetAllCategories()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCategoriesData{
		Data: categories,
	})
}
