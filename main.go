package main

import (
	initializers "awesomeProject/Initializers"
	_ "awesomeProject/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()

}

// @title      API
func main() {
	r := gin.Default()

	SetupRoutes(r)
	url := ginSwagger.URL("http://localhost:3000/docs/doc.json")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()
}
