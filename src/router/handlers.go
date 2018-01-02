package router

import (
	"github.com/gin-gonic/gin"
	"qiniupkg.com/x/log.v7"
	"net/http"
	"chrome"
	"strconv"
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

	siteType := c.DefaultQuery("site_type", "")

	strWidth := c.DefaultQuery("window_width", "1500")
	windowWidth, err := strconv.ParseUint(strWidth, 0, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ErrorCodeBadRequest,
			"error_message": err.Error(),
		})
		return
	}
	strHeight := c.DefaultQuery("window_height", "1000")
	windowHeight, err := strconv.ParseUint(strHeight, 0, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ErrorCodeBadRequest,
			"error_message": err.Error(),
		})
		return
	}

	data, err := chrome.GetScreenShot(url, siteType, windowWidth, windowHeight)
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

