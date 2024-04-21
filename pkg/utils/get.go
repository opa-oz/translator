package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"translator/pkg/config"
)

func GetConfig(c *gin.Context) (*config.Environment, error) {
	r := c.Value("Config")

	if r == nil {
		err := fmt.Errorf("could not retrieve Config")
		return nil, err
	}

	rdb, ok := r.(*config.Environment)
	if !ok {
		err := fmt.Errorf("variable Config has wrong type")
		return nil, err
	}

	return rdb, nil
}
