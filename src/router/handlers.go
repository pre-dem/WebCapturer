package router

import (
	"github.com/gin-gonic/gin"
	"qiniupkg.com/x/log.v7"
	"net/http"
	"chrome"
)

const ErrorCodeBadRequest = 400000

func GetScreenShot_v1(c *gin.Context)  {
	url, ok := c.GetQuery("url")
	if !ok {
		errorMsg := "please specify the url you want to capture"
		log.Error(errorMsg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ErrorCodeBadRequest,
			"error_message": errorMsg,
		})
		return
	}

	siteType, _ := c.GetQuery("site_type")

	data, err := chrome.GetScreenShot(url, siteType)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ErrorCodeBadRequest,
			"error_message": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "image/png", data)
}

