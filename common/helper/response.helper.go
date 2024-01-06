package helper

import (
	"github.com/gin-gonic/gin"
	"time"
)

type RespAPI struct {
	ProcessTime float64     `json:"process_time"`
	Messages    interface{} `json:"messages"`
	StatusCode  int         `json:"status_code"`
}

type RespDataAPI struct {
	RespAPI
	Data interface{} `json:"data"`
}

type RespDataAPIWithPagination struct {
	RespDataAPI
	Paginate
}

func ResponseAPI(c *gin.Context, code int, message interface{}, startTime time.Time) {
	var response RespAPI
	response.ProcessTime = float64(time.Since(startTime))
	response.Messages = message
	response.StatusCode = code

	c.JSON(code, response)
}

func ResponseDataAPI(c *gin.Context, code int, message, data interface{}, startTime time.Time) {
	var response RespDataAPI
	response.ProcessTime = float64(time.Since(startTime))
	response.Messages = message
	response.StatusCode = code
	response.Data = data

	c.JSON(code, response)
}

func SuccessResponseWithPagination(c *gin.Context, code int, message, data interface{}, paginate Paginate, startTime time.Time) {
	var response RespDataAPIWithPagination
	response.ProcessTime = float64(time.Since(startTime))
	response.Messages = message
	response.StatusCode = code
	response.Data = data
	response.Paginate = paginate

	c.JSON(code, response)
}
