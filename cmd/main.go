package main

import (
	"taskflow/lib"
	"taskflow/middleware"
	"taskflow/web/backend"

	"github.com/gin-gonic/gin"
)

// init initializes the database connection and migrates the models
func init() {
	lib.ConnectToDb()
	lib.Migrate()
}

func main() {
	r := gin.Default()
	r.POST("/signup", backend.SignUp)
	r.POST("/signin", backend.SignIn)
	r.POST("/add-item", middleware.AuthMiddleware, backend.AddTodoItem)
	r.DELETE("/delete-item/:id", middleware.AuthMiddleware, backend.DeleteTodoItem)
	
	r.GET("/", middleware.AuthMiddleware, backend.GetTodoItems)
	r.GET("/signin", backend.SignInPageHandler)
	r.GET("/signup", backend.SignUpPageHandler)
	r.GET("/signout", backend.SignOut)
	
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
