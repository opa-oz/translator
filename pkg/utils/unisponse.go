package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unisponse(c *gin.Context, text string) {
	c.JSON(http.StatusOK, gin.H{
		"text": text,
	})
}
