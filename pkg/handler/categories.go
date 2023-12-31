package handler

import (
	"marketplace"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

type getAllCategoriesData struct {
	Data []marketplace.CategoriesList `json:"data"`
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.Categories.GetAllCategories()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, getAllCategoriesData{
		Data: categories,
	})
}

func (h *Handler) getCategoriesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	category, err := h.services.Categories.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) createCategories(c *gin.Context) {
	userStatus, _ := checkStatus(c)
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
		return
	}

	var category marketplace.CategoriesList
	if err := c.BindJSON(&category); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	if _, err := h.services.Categories.GetByString(category.Title); err == nil {
		NewErrorResponse(c, http.StatusBadRequest, "already exist")
		return
	}

	id, err := h.services.Categories.Create(category)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "something went wrong")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	userStatus, _ := checkStatus(c)
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Categories.Delete(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "this brand not exists")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateCategory(c *gin.Context) {
	userStatus, _ := checkStatus(c)
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input marketplace.CategoriesList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Categories.Update(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) createCategoriesImage(c *gin.Context) {
	userStatus, _ := checkStatus(c)
	if !userStatus {
		NewErrorResponse(c, http.StatusForbidden, "you do not have sufficient rights")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	thumb, err := imageupload.ThumbnailPNG(img, 500, 500)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Categories.CreateImage(id, thumb)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getCategoriesImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	image, err := h.services.Categories.GetImage(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.File(image)

}
