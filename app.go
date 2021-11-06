package main

import (
	"core/core"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/ip", requestHandler)
	err := engine.Run(":8081")
	if err != nil {
		panic(err)
	}
}

func requestHandler(ctx *gin.Context) {
	request := ctx.Request
	url := core.GetRequestURL(request)
	data := &core.ResponseData{
		Proto:      request.Proto,
		Uri:        request.RequestURI,
		RealIp:     core.ExtractRealIp(request),
		RequestURL: url,
		Method:     request.Method,
	}
	ctx.JSON(200, data)
}
