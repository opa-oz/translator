package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// Healz godoc
// @Summary healz
// @Schemes
// @Description Check health endpoint
// @Tags utils
// @Accept json
// @Produce json
// @Success 200 {object} api.Healz.response
// @Router /healz [get]
func Healz(c *gin.Context) {
	type response struct {
		Message string `json:"message" swaggertype:"string" example:"OK"`
	}

	c.JSON(http.StatusOK, response{Message: "OK"})
}
