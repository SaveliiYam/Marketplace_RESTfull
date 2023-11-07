package handler

import (
	"marketplace"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllBasket struct {
	Data []marketplace.BusketList `json:"data"`
}

func (h *Handler) getBasket(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "have not user")
		return
	}

	basket, err := h.services.Basket.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBasket{
		Data: basket,
	})
}

func (h *Handler) createBasket(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "have not user")
		return
	}

	var input marketplace.BusketList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Basket.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
