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
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, getAllBrandsData{
		Data: brands,
	})
}

func (h *Handler) getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	brand, err := h.services.Brands.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, brand)
}

func (h *Handler) createBrand(c *gin.Context) {
	userStatus, err := checkStatus(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong!")
		return
	}
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
	}

	var brand marketplace.BrandsList
	if err := c.BindJSON(&brand); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.services.Brands.GetByString(brand.Title); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "already exist")
		return
	}

	id, err := h.services.Brands.Create(brand)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteBrand(c *gin.Context) {
	userStatus, err := checkStatus(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong!")
		return
	}
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Brands.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateBrand(c *gin.Context) {
	userStatus, err := checkStatus(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong!")
		return
	}
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input marketplace.BrandsList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Brands.Update(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
