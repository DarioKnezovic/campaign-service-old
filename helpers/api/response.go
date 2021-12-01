package api

import "github.com/gin-gonic/gin"

//ResponseData structure
type ResponseData struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

//Message returns map data
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func ResponseWithData(c *gin.Context, status int, data map[string]interface{}) {
	c.JSON(status, data)
}