package routers

import (
	"github.com/gin-gonic/gin"
	v1Controllers "github.com/DarioKnezovic/campaign-service/controllers/v1"
	"github.com/DarioKnezovic/campaign-service/middlewares"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	v1 := r.Group("/api/v1")

	v1.Use(middlewares.AuthMiddleware())
	{
		v1.GET("test", v1Controllers.TestHome)
	}

	return r
}
