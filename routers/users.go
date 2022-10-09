package routers

import (
	"gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(route *gin.Engine) {
	//added new
	route.GET("/v1/users/", controllers.GetUserData)
	route.POST("/v1/users/", controllers.Create)

	//Add All route
	//TestRoutes(route)
}
