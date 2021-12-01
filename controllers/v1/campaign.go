package v1

import (
	"github.com/gin-gonic/gin"
	v1Services "github.com/DarioKnezovic/campaign-service/services/v1"
	r "github.com/DarioKnezovic/campaign-service/helpers/api"
)

func TestHome(c *gin.Context)  {

	testData := v1Services.ReturnTestData()

	r.ResponseWithData(c, r.STATUS_OK, testData)
}