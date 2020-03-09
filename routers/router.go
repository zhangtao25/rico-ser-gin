package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine  {
	r := gin.New()

	r.GET("/job", GetJobs)

	return r
}