package handler

import (
	"marketplace"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	userStatus, err := checkStatus(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong!")
		return
	}
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
	}

	var input marketplace.ProductList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	brandId, err := h.services.Brands.GetByString(input.Brand)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "have not this brand") // Если нет указанного бренда
		return
	}
	categoriesId, err := h.services.Categories.GetByString(input.Category)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "have not this category") // Если нет указанной категории
		return
	}

	id, err := h.services.Products.Create(input, brandId, categoriesId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllProductsData struct {
	Data []marketplace.ProductList `json:"data"`
}

func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.services.Products.GetAll()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, getAllProductsData{
		Data: products,
	})
}

func (h *Handler) getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	product, err := h.services.Products.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) deleteProduct(c *gin.Context) {
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

	err = h.services.Products.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateProduct(c *gin.Context) {
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

	var input marketplace.ProductList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	brandId, err := h.services.Brands.GetByString(input.Brand)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "have not this brand") // Если нет указанного бренда
		return
	}
	categoriesId, err := h.services.Categories.GetByString(input.Category)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "have not this category") // Если нет указанной категории
		return
	}

	if err := h.services.Products.Update(id, brandId, categoriesId, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
