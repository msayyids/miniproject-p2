package main

import (
	"miniproject/config"
	"miniproject/controller"
	"miniproject/middleware"
	"miniproject/repo"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	"github.com/swaggo/gin-swagger"

	_ "miniproject/docs"
)

// @title hotels API
// @version 1.0
// @description This is a sample hotel server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2

func main() {
	db := config.InitDb()
	repo := repo.Repo{DB: db}
	c := controller.Controllers{Controller: repo}
	a := middleware.Auth{Authentication: repo}

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/register", c.Register)
	r.POST("/login", c.Login)

	homepage := r.Group("/hocation")
	auth := homepage.Use(a.AuthUsers())

	auth.PUT("/topup", c.EditAmount)
	auth.POST("/user", c.GetLoggedInUserInfo)

	auth.GET("/rooms", c.FindAvailableRoom)
	auth.GET("/booking", c.GetBooking)
	auth.POST("/booking", c.CreateBooking)
	auth.PUT("/booking/:id", c.UpdateBooking)

	r.Run(":8080")
}
