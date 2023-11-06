package handler

import (
	"marketplace"
	"net/http"
	"strconv"

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

func (h *Handler) getBrandById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	brand, err := h.services.Brands.GetBrandById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, brand)
}
