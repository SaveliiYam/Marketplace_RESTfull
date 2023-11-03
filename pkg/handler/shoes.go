package handler

import (
	"marketplace"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllShoesData struct {
	Data []marketplace.ProductList `json:"data"`
}

func (h *Handler) getAllShoes(c *gin.Context) {
	shoes, err := h.services.Shoes.GetAllShoes()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllShoesData{
		Data: shoes,
	})
}
