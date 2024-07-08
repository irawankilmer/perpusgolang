package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/perpustakaanonline/controllers"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", controllers.RegisterUser)
		}
	}
}
