package controllers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Raiz retorna links para os controllers
// @Summary Raiz
// @Description Retorna links para os controllers
// @Tags Raiz
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func (ct *Controller) Raiz(c *gin.Context) {
	addresses := make(map[string]string)
	val := reflect.ValueOf(ct.controllerNames).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i).String()
		addresses[field] = fmt.Sprintf("%s/api/v1/%s", ct.config.ServerHost, field)
	}
	c.JSON(http.StatusOK, addresses)
}
