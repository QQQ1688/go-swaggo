package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"

	_ "go-api/docs" // import 檔案夾的名稱要以 go.mod 的名稱來命名，再加檔案夾名稱

	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

// sets up an association in which GetDatass handles requests to the /mysql endpoint path.
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server mysql server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	// Initialize a Gin router using Default.
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		iplogs := v1.Group("/iplogs")
		{
			iplogs.GET("", controller.GetAllDatas)
			iplogs.GET("/:ip", controller.GetDataByIP)
			iplogs.POST("", controller.PostData)
		}
	}
	// Use the Run function to attach the router to an http.Server and start the server.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
