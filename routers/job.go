package routers

import (
	"github.com/gin-gonic/gin"
	"rico-ser-gin/pkg/e"

	//"rico-ser-gin/pkg/e"
	"rico-ser-gin/pkg/app"
	"net/http"
)

func GetJobs(c *gin.Context)  {
	appG := app.Gin{C: c}

	type job struct {
		id string
		name string
	}

	//Job := job{id:"1",name:"zt"}
	appG.Response(http.StatusOK, e.SUCCESS,"ssss")
}