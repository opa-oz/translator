package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "translator/cmd/docs"
	"translator/pkg/api"
	"translator/pkg/config"
	"translator/pkg/middlewares"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.Production {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	r.Use(middlewares.RequestLogger())
	r.Use(middlewares.ResponseLogger())

	r.GET("/healz", api.Healz)
	r.GET("/ready", api.Ready)
	// {@link https://github.com/swaggo/gin-swagger?tab=readme-ov-file}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	rg := r.Group("/api")
	rg.Use(middlewares.CfgMiddleware(cfg))

	{
		rg.POST("/translate/:fromLang/:toLang", api.Translate)
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	err = r.Run(port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
