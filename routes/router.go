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
		routerV1.POST("category/add", v1.AddCate)
		routerV1.GET("categories", v1.ListCate)
		routerV1.PUT("category/:id", v1.UpdateCate)
		routerV1.DELETE("category/:id", v1.DeleteCate)
		// Article router
		routerV1.POST("article/add", v1.AddArt)
		routerV1.GET("articles", v1.ListArts)
		routerV1.GET("article/category/:id", v1.GetArtForCate)
		routerV1.GET("article/info/:id", v1.GetArtInfo)
		routerV1.PUT("article/:id", v1.UpdateArt)
		routerV1.DELETE("article/:id", v1.DeleteArt)
	}

	r.Run(utils.HttpPort)
}
