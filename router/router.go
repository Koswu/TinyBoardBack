package router

import (
	"boarderbackend/middleware/cors"
	"boarderbackend/middleware/jwt"
	"boarderbackend/pkgs/setting"
	v1 "boarderbackend/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	if setting.IsDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(cors.AllowCORS())
	authApi := r.Group("/auth")
	{
		authApi.POST("/login", GetAuth)
		authApi.POST("/register", RegisterUser)
	}

	apiv1 := r.Group("/api/v1")
	{
		readApi := apiv1.Group("/read")
		{
			readApi.GET("/comment", v1.GetComments)
			readApi.GET("/comment/:id", v1.GetComment)
		}
		writeApi := apiv1.Group("/write")
		writeApi.Use(jwt.JWT())
		{
			writeApi.POST("/comment", v1.PutComment)
		}
	}
	return r

}