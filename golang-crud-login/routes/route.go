package routes

import (
	"golang-crud-login/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// load templates
	r.LoadHTMLGlob("views/*.html")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// View routes
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "layout.html", gin.H{"title": "Login", "content": "login"})
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "layout.html", gin.H{"title": "Register", "content": "register"})
	})
	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.Redirect(302, "/login")
			return
		}
		c.HTML(200, "layout.html", gin.H{"title": "Dashboard", "username": user})
	})

	// API routes
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
