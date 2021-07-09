package main

import (
	"github.com/ddovale/short-url/src/api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := getRouter()
	defineRoutes(engine)
	engine.Run()
}

func getRouter() *gin.Engine {
	return gin.Default()
}

func defineRoutes(engine *gin.Engine) {
	engine.GET("/ping", controllers.PingHandler)

	engine.GET("/urls", controllers.GetAllUrls)
	engine.GET("/urls/:id", controllers.GetById)
	engine.POST("/urls", controllers.CreateUrl)
}
