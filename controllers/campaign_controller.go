package v1

import (
	"github.com/gin-gonic/gin"
)

func TestHome(c *gin.Context)  {

	c.JSON(200, gin.H{
		"message": "Hello you",
	})
}