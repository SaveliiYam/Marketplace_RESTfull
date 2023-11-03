package handler

import (
	"marketplace"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllBrandsData struct {
	Data []marketplace.BrandsList `json:"data"`
}

func (h *Handler) getAllBrands(c *gin.Context) {
	brands, err := h.services.Brands.GetAllBrands()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBrandsData{
		Data: brands,
	})
}
