package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Raiz retorna links para os controllers
// @Summary Raiz
// @Description Retorna links para os controllers
// @Tags Raiz
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router / [get]
func (ct *Controller) Raiz(c *gin.Context) {
	addresses := make(map[string]string)
	for _, address := range ct.controllerNames {
		addresses[address] = fmt.Sprintf("%s/api/v1/%s", ct.config.ServerHost, address)
	}
	c.JSON(http.StatusOK, addresses)
}
