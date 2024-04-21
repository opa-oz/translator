package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// Ready godoc
// @Summary ready
// @Schemes
// @Description Check readiness
// @Tags utils
// @Accept json
// @Produce json
// @Success 200 {object} api.Ready.response
// @Router /ready [get]
func Ready(c *gin.Context) {
	type response struct {
		Message string `json:"message" swaggertype:"string" example:"Ready"`
	}

	c.JSON(http.StatusOK, response{Message: "Ready"})
}
