package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cristhianfdx/order-service/internal/ports"
)

type OrderHandler struct {
	service ports.OrderService
}

func NewOrderHandler(s ports.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		Item   string  `json:"item"`
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.CreateOrder(req.Item, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	}
	c.JSON(http.StatusOK, order)
}
