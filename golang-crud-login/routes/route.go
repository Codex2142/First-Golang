package routes

import (
	"golang-crud-login/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	product := r.Group("/products")
	{
		product.GET("/", controllers.GetProducts)
		product.POST("/", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	return r
}
