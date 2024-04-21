package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unirror(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err,
	})
}
