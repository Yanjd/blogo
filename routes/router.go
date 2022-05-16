package routes

import (
	v1 "blogo/api/v1"
	"blogo/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// User router
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.ListUsers)
		routerV1.PUT("user/:id", v1.UpdateUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		// Category router

		// Article router
	}

	r.Run(utils.HttpPort)
}
