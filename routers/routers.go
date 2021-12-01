package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/DarioKnezovic/campaign-service/controllers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("test", controllers.TestHome)

	return r
}
